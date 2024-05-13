package handlers

import (
	"fmt"
	"net/http"
	"time"
)

var serverStartTime time.Time

func NewResponseWriter(w http.ResponseWriter) {
	panic("unimplemented")
}

func init() {
	serverStartTime = time.Now()
}

func ServerAlive(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(serverStartTime)
	fmt.Fprintln(w, "Server is alive since : ", uptime.String())
}

