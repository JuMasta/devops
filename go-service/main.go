package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


var requestCount = promauto.NewCounter(prometheus.CounterOpts{
	Name: "go_app_requests_count",
	Help: "Total App HTTP Requests",
})


func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my awesome website! " + r.Host)
	requestCount.Inc()
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(" health check request\n")
	fmt.Printf("application is healthy\n")
	w.WriteHeader(http.StatusOK)
    w.Write([]byte("200 - Everything is healthy!"))
}

func readinessCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(" readiness check request\n")
	fmt.Printf("application is ready\n")
	w.WriteHeader(http.StatusOK)
    w.Write([]byte("200 - Everything is ready!"))
}

func main() {

	wg := new(sync.WaitGroup)
	wg.Add(2)

	http.HandleFunc("/readiness", readinessCheck)
	http.HandleFunc("/healthz", healthCheck)
	http.HandleFunc("/", getRoot)

	go func() {
		http.ListenAndServe(":8080", nil)
		wg.Done()
	}()

	go func() {
		http.ListenAndServe(":8081", promhttp.Handler())
		wg.Done()
	}()

	wg.Wait()

}
