package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Create a counter metric
var requestCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests received.",
	})

// Create a gauge metric
var temperatureGauge = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "room_temperature_celsius",
		Help: "Current room temperature in Celsius.",
	})

func init() {
	// Register the metrics with Prometheus
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(temperatureGauge)
}

func main() {
	// Simulate temperature changes
	go func() {
		for {
			temperatureGauge.Set(20 + rand.Float64()*10) // Random temperature between 20 and 30
			time.Sleep(5 * time.Second)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestCounter.Inc() // Increment the request counter
		w.Write([]byte("Hello, Prometheus!"))
	})

	// Expose the registered metrics at `/metrics` endpoint
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Starting server on :8085")
	log.Fatal(http.ListenAndServe(":8085", nil))
}
