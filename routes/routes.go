package routes

import (
	"net/http"

	"github.com/ECecillo/GoProm/handlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func AddRoutes(mux *http.ServeMux) {
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("GET /api/liveliness", handlers.ServerAlive)
}
