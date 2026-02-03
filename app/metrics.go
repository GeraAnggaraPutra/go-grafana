package app

import "github.com/prometheus/client_golang/prometheus"

var (
	customCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_custom_request_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "status"},
	)

	customErrorCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_custom_error_total",
			Help: "Total number of errors returned by the API",
		},
		[]string{"path", "status"},
	)
)

func init() {
	prometheus.MustRegister(
		customCounter,
		customErrorCounter,
	)
}
