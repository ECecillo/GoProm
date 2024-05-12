package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var serverStartTime time.Time
var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "Number of request handled by Ping Handler",
	},
)

func init() {
	serverStartTime = time.Now()
}

func ServerAlive(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(serverStartTime)
	fmt.Fprintln(w, "Server is alive since : ", uptime.String())
}

func ping(w http.ResponseWriter, r *http.Request) {
	pingCounter.Inc()
	fmt.Fprintln(w, "PONG")
}

func main() {
	PORT := flag.String("PORT", ":3000", "Exposed server port")
	router := http.NewServeMux()
	server := &http.Server{
		Addr:    *PORT,
		Handler: router,
	}
	prometheus.MustRegister(pingCounter)

	router.HandleFunc("GET /api/liveliness", ServerAlive)
	router.HandleFunc("GET /api/ping", ping)
	router.Handle("/metrics", promhttp.Handler())

	fmt.Println("Server running")
	server.ListenAndServe()
}
