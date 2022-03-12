package metrics

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func Start(port uint) {
	address := fmt.Sprintf("0.0.0.0:%d", port)
	log.Println("Launching Prometheus Metric exporter on: " + address + "/metrics")
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(address, nil); err != nil {
			log.Error("Failed to start Prometheus Metrics Exporter - " + address)
		}
	}()
}
