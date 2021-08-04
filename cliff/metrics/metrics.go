package metrics

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func Export() {
	log.Println("Launching Prometheus Metric exporter on: 0.0.0.0:2112/metrics")
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe("0.0.0.0:2112", nil)
	}()
}
