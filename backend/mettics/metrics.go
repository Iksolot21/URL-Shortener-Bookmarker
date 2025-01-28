package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promauto"
)
var (
 HttpRequestsTotal = promauto.NewCounter(prometheus.CounterOpts{
   Name: "http_requests_total",
   Help: "Total number of http requests",
})
)