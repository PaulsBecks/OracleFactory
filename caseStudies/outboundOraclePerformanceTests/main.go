package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type PerformanceTest struct {
	outputFileName  string
	oracleEndpoint  string
	keyVariableName string
	body            string
}

type PerformanceTestRun struct {
	totalEvents int
	waitGroup   sync.WaitGroup
	totalStart  time.Time
	latencies   []EventMeasurement
	test        PerformanceTest
	guard       chan struct{}
	mu          *sync.Mutex
	server      http.Server
	events      map[int]EventMeasurement
}

func NewPerformanceTestRun(performanceTest *PerformanceTest, maxEventsParallel int) *PerformanceTestRun {
	return &PerformanceTestRun{guard: make(chan struct{}, maxEventsParallel), mu: &sync.Mutex{}, totalEvents: 10, test: *performanceTest, events: make(map[int]EventMeasurement)}
}

type EventMeasurement struct {
	start    time.Time
	latency  float64
	success  bool
	workerID int
}

var (
	currentPerformaceTestRun *PerformanceTestRun
	server                   http.Server
)

func writeToCSV(line []string, file *os.File) {
	writer := csv.NewWriter(file)
	defer writer.Flush()

	err := writer.Write(line)
	if err != nil {
		log.Fatalf("Unable to write to %s", file.Name())
	}
}

func sendRequestToInboundOracle(endpoint string, body []byte) {
	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	// TODO: check for http response code
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
}
func (p *PerformanceTestRun) timeEvent(worker int) {
	var event map[string]interface{}
	json.Unmarshal([]byte(p.test.body), &event)
	event[p.test.keyVariableName] = worker

	jsonStr, _ := json.Marshal(event)
	sendRequestToInboundOracle(p.test.oracleEndpoint, jsonStr)

	start := time.Now()
	p.events[worker] = EventMeasurement{
		start:    start,
		workerID: worker,
	}
	fmt.Println(fmt.Sprintf("New event created %d", worker))
}

func (p *PerformanceTestRun) start() (averageLatency, throughputPerSecond float64) {
	defer p.waitGroup.Wait()
	p.totalStart = time.Now()
	for i := 0; i < p.totalEvents; i++ {
		p.waitGroup.Add(1)
		p.guard <- struct{}{}
		go p.timeEvent(i)
	}
	p.waitGroup.Wait()
	totalElapsed := time.Since(p.totalStart)
	avgLatency, _ := computeAverageLatency(p.latencies)
	return avgLatency, float64(p.totalEvents) / totalElapsed.Seconds()
}

func (p *PerformanceTest) runAll(repetitions int) {
	log.Printf("Execute performance test for endpoint %s.", p.oracleEndpoint)
	var err error
	file, err := os.Create(p.outputFileName)
	if err != nil {
		log.Fatalf("Unable to open %s", p.outputFileName)
		return
	}
	writeToCSV([]string{"latency", "throughput", "parallel events"}, file)
	defer file.Close()
	for i := 0; i < repetitions; i++ {
		for _, maxParallel := range []int{1, 2, 3, 4, 5, 10, 20} { //, 30, 40, 50, 100} {
			log.Printf("Maximum of events created in parallel %d", maxParallel)
			performanceTestRun := NewPerformanceTestRun(p, maxParallel)
			currentPerformaceTestRun = performanceTestRun
			avgLatency, throughput := performanceTestRun.start()
			writeToCSV([]string{fmt.Sprintf("%f", avgLatency), fmt.Sprintf("%f", throughput), fmt.Sprintf("%d", maxParallel)}, file)
		}
	}
	stopServer()
}

func (p *PerformanceTestRun) handler(w http.ResponseWriter, r *http.Request) {
	stop := time.Now()

	var event map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		fmt.Printf("Error occured %v", err)
		return
	}
	workerID := event[p.test.keyVariableName].(string)
	intWorkerID, err := strconv.Atoi(workerID)
	if err != nil {
		fmt.Printf("Error occured %v", err)
		return
	}
	fmt.Println(workerID, event)
	measurement := p.events[intWorkerID]
	elapsed := stop.Sub(measurement.start)
	measurement.latency = float64(elapsed)
	p.mu.Lock()
	p.latencies = append(p.latencies, measurement)
	p.mu.Unlock()
	p.waitGroup.Done()
	<-p.guard
}

func startServer() {
	server = http.Server{Addr: ":7890"}
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			currentPerformaceTestRun.handler(w, r)
		})
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Server closed with error %v", err)
		}
	}()
}

func stopServer() {
	server.Shutdown(context.Background())
	server.Close()
}

func computeAverageLatency(eventMeasurements []EventMeasurement) (float64, error) {
	amount := len(eventMeasurements)
	if amount <= 0 {
		return 0, fmt.Errorf("can't calculate average of 0 elements")
	}
	total := 0.
	for _, measurement := range eventMeasurements {
		total += measurement.latency
	}
	return total / float64(amount), nil
}

func main() {
	repetitions := 10
	/*hyperledgerCreateAssetTest := &PerformanceTest{
		outputFileName: "hyperledgerCreateAssetTest.csv",
		oracleEndpoint: "http://localhost:8080/webServiceListeners/1/events",
		body:           `{"assetID":"1","color":"green", "size":"m", "owner":"me", "appraisedValue":"1k"}`,
	}
	hyperledgerCreateAssetTest.runAll(repetitions)*/
	startServer()
	ethereumMintTokenTest := &PerformanceTest{
		outputFileName:  "ethereumTransferTokenTest.csv",
		oracleEndpoint:  "http://localhost:8080/webServiceListeners/3/events",
		body:            `{"receiver":"0x40536521353F9f4120A589C9ddDEB6188EF61922","amount":0}`,
		keyVariableName: "amount",
	}

	sendRequestToInboundOracle("http://localhost:8080/webServiceListeners/2/events", []byte(`{"receiver":"0x40536521353F9f4120A589C9ddDEB6188EF61922","amount":100000}`))

	ethereumMintTokenTest.runAll(repetitions)
}
