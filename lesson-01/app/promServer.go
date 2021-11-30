package app

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PromServer struct {
	Router  *http.ServeMux
	Metrics *AppMetrics
}

// CreatePromServer function
func CreatePromServer() (ps *PromServer, err error) {
	ps = new(PromServer)

	ps.Router = http.NewServeMux()
	ps.Router.Handle("/metrics", promhttp.Handler())

	ps.Metrics, err = CreateAppMetrics()
	return
}
