package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const port int = 4200

var aliveMsg = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "alive_message",
	Help: "Just to see, if service is still running",
})

func Run() {
	fmt.Printf("Running Server on port: %d", port)
	r := mux.NewRouter()

	prometheus.MustRegister(aliveMsg)
	aliveMsg.Add(69)
	r.HandleFunc("/", handleAlive)
	r.Handle("/metrics", promhttp.Handler())

	portSting := fmt.Sprintf(":%d", port)
	http.ListenAndServe(portSting, r)
}
