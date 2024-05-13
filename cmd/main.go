package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/ECecillo/GoProm/handlers"
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


func init() {
	prometheus.Register(middleware.TotalRequests)
}


func main() {
	PORT := flag.String("PORT", ":8080", "Exposed server port")
	router := http.NewServeMux()
	middelwares := createStack(middleware.Prometheus)
	server := &http.Server{
		Addr:    *PORT,
		Handler: middelwares(router),
	}

	router.HandleFunc("GET /api/liveliness", handlers.ServerAlive)
	router.Handle("/metrics", promhttp.Handler())

	fmt.Println("Server running")
	server.ListenAndServe()
}
