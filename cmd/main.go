package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/ECecillo/GoProm/middleware"
	"github.com/ECecillo/GoProm/types"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func createStack(xs ...types.Middleware) types.Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}
		return next
	}
}

var serverStartTime time.Time

func NewResponseWriter(w http.ResponseWriter) {
	panic("unimplemented")
}

func init() {
	serverStartTime = time.Now()
	prometheus.Register(middleware.TotalRequests)
}

func ServerAlive(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(serverStartTime)
	fmt.Fprintln(w, "Server is alive since : ", uptime.String())
}

func main() {
	PORT := flag.String("PORT", ":8080", "Exposed server port")
	router := http.NewServeMux()
	middelwares := createStack(middleware.Prometheus)
	server := &http.Server{
		Addr:    *PORT,
		Handler: middelwares(router),
	}

	router.HandleFunc("GET /api/liveliness", ServerAlive)
	router.Handle("/metrics", promhttp.Handler())

	fmt.Println("Server running")
	server.ListenAndServe()
}
