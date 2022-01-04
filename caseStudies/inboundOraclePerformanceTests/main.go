package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type PerformanceTest struct {
	outputFileName string
	oracleEndpoint string
	body           string
}

type PerformanceTestRun struct {
	totalEvents int
	waitGroup   sync.WaitGroup
	totalStart  time.Time
	latencies   []EventMeasurement
	test        PerformanceTest
	guard       chan struct{}
	mu          *sync.Mutex
}

func NewPerformanceTestRun(performanceTest *PerformanceTest, maxEventsParallel int) *PerformanceTestRun {
	return &PerformanceTestRun{guard: make(chan struct{}, maxEventsParallel), mu: &sync.Mutex{}, totalEvents: 20, test: *performanceTest}
}

type EventMeasurement struct {
	latency  float64
	success  bool
	workerID int
}

func writeToCSV(line []string, file *os.File) {
	writer := csv.NewWriter(file)
	defer writer.Flush()

	err := writer.Write(line)
	if err != nil {
		log.Fatalf("Unable to write to %s", file.Name())
	}
}

func (p *PerformanceTestRun) timeEvent(worker int) {
	var jsonStr = []byte(p.test.body)
	req, _ := http.NewRequest("POST", p.test.oracleEndpoint, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	start := time.Now()
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
	elapsed := time.Since(start)
	measurement := EventMeasurement{latency: elapsed.Seconds(), success: true, workerID: worker}
	fmt.Println(start.String(), worker)
	p.mu.Lock()
	p.latencies = append(p.latencies, measurement)
	p.mu.Unlock()
	p.waitGroup.Done()
	<-p.guard
}

func (p *PerformanceTestRun) start() (averageLatency, throughputPerSecond float64) {
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
		for _, maxParallel := range []int{1, 2, 3, 4} { //, 30, 40, 50, 100} {
			log.Printf("Maximum of events created in parallel %d", maxParallel)
			performanceTestRun := NewPerformanceTestRun(p, maxParallel)
			avgLatency, throughput := performanceTestRun.start()
			writeToCSV([]string{fmt.Sprintf("%f", avgLatency), fmt.Sprintf("%f", throughput), fmt.Sprintf("%d", maxParallel)}, file)
		}
	}
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

const BASE_URL = "http://localhost:8080/"

func subscribe(outboundOracleID int, smartContractAddress string) {
	params := map[string]interface{}{
		"Token":                "",
		"Topic":                "test-topic",
		"Filter":               "",
		"Callback":             "Callback",
		"SmartContractAddress": smartContractAddress,
	}
	json, _ := json.Marshal(params)
	req, err := http.NewRequest("POST", fmt.Sprintf("%soutboundOracles/%d/subscribe", BASE_URL, outboundOracleID), bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	fmt.Println(resp)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
}

func unsubscribe(outboundOracleID int, smartContractAddress string) {
	params := map[string]interface{}{
		"Token": "",
		"Topic": "test-topic",
	}
	json, _ := json.Marshal(params)
	http.NewRequest("POST", fmt.Sprintf("%soutboundOracles/%d/subscribe", BASE_URL, outboundOracleID), bytes.NewBuffer(json))
}

func main() {
	repetitions := 5
	now := time.Now().UTC()

	// subscribe smart contract to hyperledger provider
	subscribe(2, "test-contract")
	hyperledgerCreateAssetTest := &PerformanceTest{
		outputFileName: fmt.Sprintf("hyperledger1Subscription%s_.csv", now),
		oracleEndpoint: BASE_URL + "providers/1/events",
		body:           `{"number":1}`,
	}
	hyperledgerCreateAssetTest.runAll(repetitions)

	// subscribe smart contract to hyperledger provider
	subscribe(2, "test-contract2")
	hyperledgerCreateAssetTest.outputFileName = fmt.Sprintf("hyperledger2Subscription%s_.csv", now)
	hyperledgerCreateAssetTest.runAll(repetitions)

	// subscribe smart contract to hyperledger provider
	subscribe(2, "test-contract3")
	hyperledgerCreateAssetTest.outputFileName = fmt.Sprintf("hyperledger3Subscription%s_.csv", now)
	hyperledgerCreateAssetTest.runAll(repetitions)

	/*ethereumMintTokenTest := &PerformanceTest{
		outputFileName: "ethereumMintTokenTest.csv",
		oracleEndpoint: "http://localhost:8080/providers/2/events",
		body:           `{"receiver":"0x40536521353F9f4120A589C9ddDEB6188EF61922","amount":100}`,
	}
	ethereumMintTokenTest.runAll(repetitions)

	ethereumTransferTokenTest := &PerformanceTest{
		outputFileName: "ethereumTransferTokenTest.csv",
		oracleEndpoint: "http://localhost:8080/providers/3/events",
		body:           `{"receiver":"0x40536521353F9f4120A589C9ddDEB6188EF61922","amount":1}`,
	}
	ethereumTransferTokenTest.runAll(repetitions)*/
}
