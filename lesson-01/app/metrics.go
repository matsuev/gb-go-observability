package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const Namespace = "myservice"

type AppMetrics struct {
	totalRequests         prometheus.Counter
	urlRequests           *prometheus.CounterVec
	responseTimeHistogram *prometheus.HistogramVec
}

func CreateAppMetrics() (am *AppMetrics, err error) {
	defer func() {
		if r := recover(); r != nil {
			am = nil
			err = fmt.Errorf("error: CreateAppMetrics: %v", r)
		}
	}()

	am = new(AppMetrics)

	am.totalRequests = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: Namespace,
		Name:      "total_requests",
		Help:      "Total number of requests",
	})
	prometheus.MustRegister(am.totalRequests)

	am.urlRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: Namespace,
			Name:      "url_requests",
			Help:      "Number of request per route",
		},
		[]string{
			"uri",
		},
	)
	prometheus.MustRegister(am.urlRequests)

	am.responseTimeHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: Namespace,
			Name:      "response_time_histogram",
			Help:      "Response time histogram",
			Buckets:   prometheus.LinearBuckets(.100, .1, 10),
		},
		[]string{
			"uri",
		},
	)
	prometheus.MustRegister(am.responseTimeHistogram)

	return
}

// Collect function
func (am *AppMetrics) Collect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		am.totalRequests.Inc()
		am.urlRequests.WithLabelValues(r.URL.Path).Inc()

		startTime := time.Now()
		next.ServeHTTP(w, r)
		endTime := time.Since(startTime)

		am.responseTimeHistogram.WithLabelValues(r.URL.Path).Observe(endTime.Seconds())
	})
}
