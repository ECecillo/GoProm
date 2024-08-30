package main

import (
	"context"
	"flag"

	"github.com/ECecillo/GoProm/middleware"
	"github.com/ECecillo/GoProm/server"
	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	prometheus.Register(middleware.TotalRequests)
}

func main() {
	PORT := flag.String("PORT", ":9000", "Exposed server port")
	config := server.Config{
		Host: "localhost",
		Port: *PORT,
	}
	ctx := context.Context{}


	err := server.Create(ctx, &config)
	if err != nil {
		panic(err)
	}
}
