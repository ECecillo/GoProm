package middleware

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

// Pas convaincu que ce soit la bonne manière de faire
var TotalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of request handled",
	},
	[]string{"path"},
)

func Prometheus(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		TotalRequests.WithLabelValues(r.RequestURI)
	})
}
