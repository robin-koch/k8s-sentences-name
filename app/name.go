package main

import (
    "fmt"
    "log"
	"net/http"
	"math/rand"
	"time"
	"github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    httpReqs = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "sentence_requests_total",
			Help: "Number of requests",
		},
		[]string{"type"},
	)
)

func GetName() (string) {
	names := [5]string{"Graham", "John", "Terry", "Eric", "Michael", "Robin", "Martin"}

	rand.Seed(time.Now().UnixNano())
	return names[rand.Intn(len(names))]
}

func handler(httpReqs *prometheus.CounterVec) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        name := GetName()

		fmt.Fprintf(w, "%s", name)

        m := httpReqs.WithLabelValues("name")
        m.Inc()
    }
}

func main() {
	prometheus.MustRegister(httpReqs)

    http.HandleFunc("/", handler(httpReqs))
    http.Handle("/metrics", promhttp.Handler())
    log.Fatal(http.ListenAndServe(":8080", nil))
}
