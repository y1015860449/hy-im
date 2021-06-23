package server

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func StartPrometheus(addr string) (err error) {
	r := "8085"
	if addr != "" {
		r = addr
	}
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		if err := http.ListenAndServe(r, nil); err != nil {
			log.Fatalf("start prometheus err (%v)", err)
		}
	}()
	return
}
