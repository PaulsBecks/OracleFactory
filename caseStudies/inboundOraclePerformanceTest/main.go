package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

const (
	fileName          = "output"
	oracleEndpoint    = "http://localhost:8080/inboundOracles/1/events"
	body              = `{"assetID":"1","color":"green", "size":"m", "owner":"me", "appraisedValue":"1k"}`
	totalEvents       = 1000
	processesParallel = 100
)

var (
	file       *os.File
	waitGroup  sync.WaitGroup
	guard      = make(chan struct{}, processesParallel)
	totalStart time.Time
	latencies  []EventMeasurement
	mu         = &sync.Mutex{}
)

type EventMeasurement struct {
	latency  float64
	success  bool
	workerID int
}

func writeToCSV(line []string) {
	writer := csv.NewWriter(file)
	defer writer.Flush()

	err := writer.Write(line)
	if err != nil {
		log.Fatalf("Unable to write to %s.csv", fileName)
	}
}

func timeIt(start time.Time, success bool, worker int) {
	elapsed := time.Since(start)
	//totalElapsed := time.Since(totalStart)
	measurement := EventMeasurement{latency: elapsed.Seconds(), success: success, workerID: worker}
	mu.Lock()
	latencies = append(latencies, measurement)
	mu.Unlock()
	//writeToCSV([]string{fmt.Sprintf("%f", elapsed.Seconds()), fmt.Sprintf("%f", totalElapsed.Seconds()), success, worker})
}

func timeEvent(worker int) {
	var jsonStr = []byte(body)
	req, _ := http.NewRequest("POST", oracleEndpoint, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	if err != nil {
		timeIt(start, false, worker)
		log.Fatal(err)
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	log.Println(string(responseBody))
	if err != nil {
		timeIt(start, false, worker)
		log.Fatal(err)
	}
	timeIt(start, true, worker)
	waitGroup.Done()
	<-guard
}

func runTest() (averageLatency, throughputPerSecond float64) {
	var err error
	file, err = os.Create(fmt.Sprintf("%s.csv", fileName))
	if err != nil {
		log.Fatalf("Unable to open %s.csv", fileName)
		return 0., 0.
	}
	defer file.Close()
	defer waitGroup.Wait()
	totalStart = time.Now()
	for i := 0; i < totalEvents; i++ {
		waitGroup.Add(1)
		guard <- struct{}{}
		go timeEvent(i)
	}
	waitGroup.Wait()
	totalElapsed := time.Since(totalStart)
	avgLatency, _ := computeAverageLatency(latencies)
	return avgLatency, totalEvents / totalElapsed.Seconds()
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
	avgLatency, throughputPerSecond := runTest()
	fmt.Printf("%f, %f", avgLatency, throughputPerSecond)
}
