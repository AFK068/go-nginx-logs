package domain

import (
	"github.com/es-debug/backend-academy-2024-go-template/pkg/timeutils"
	"github.com/streadway/quantile"
)

type LogReport struct {
	FilterConfig      *FlagConfig
	NumberRequests    int
	TotalResponseSize int
	MaxResponseSize   int
	MinResponseSize   int
	ResourceCount     map[string]int
	StatusCount       map[int]int
	QuantileEstimator *quantile.Estimator
}

func NewLogReport(fc *FlagConfig) *LogReport {
	return &LogReport{
		FilterConfig:      fc,
		MaxResponseSize:   0,
		MinResponseSize:   int(^uint(0) >> 1), // MaxInt
		ResourceCount:     make(map[string]int),
		StatusCount:       make(map[int]int),
		QuantileEstimator: quantile.New(quantile.Known(0.95, 0.001)),
	}
}

func (lr *LogReport) Update(log *NGINX) {
	// If the time of the log is not in the range, then we do not update the report.
	if !timeutils.InTimeSpan(lr.FilterConfig.From, lr.FilterConfig.To, log.TimeLocal) {
		return
	}

	// If the log does not match the filter, then we do not update the report.
	if !lr.FilterConfig.FilterMatch(log) {
		return
	}

	// Added 1 to the number of requests.
	lr.incrementRequestCount()

	// Added the size of the response to the total size of the responses.
	lr.addResponseSize(log.BodyBytesSent)

	// Updated the maximum and minimum sizes of the response.
	lr.updateMaxResponseSizeStats(log.BodyBytesSent)
	lr.updateMinResponseSizeStats(log.BodyBytesSent)

	// Updated the number of requests for the resource and the status.
	lr.incrementResourceCount(log.Request)
	lr.incrementStatusCount(log.Status)

	// Added the size of the response to the quantile estimator.
	lr.addToQuantileEstimator(log.BodyBytesSent)
}

func (lr *LogReport) incrementRequestCount() {
	lr.NumberRequests++
}

func (lr *LogReport) addResponseSize(bodyBytesSent int) {
	lr.TotalResponseSize += bodyBytesSent
}

func (lr *LogReport) updateMaxResponseSizeStats(bodyBytesSent int) {
	lr.MaxResponseSize = max(lr.MaxResponseSize, bodyBytesSent)
}

func (lr *LogReport) updateMinResponseSizeStats(bodyBytesSent int) {
	lr.MinResponseSize = min(lr.MinResponseSize, bodyBytesSent)
}

func (lr *LogReport) incrementResourceCount(request string) {
	lr.ResourceCount[request]++
}

func (lr *LogReport) incrementStatusCount(status int) {
	lr.StatusCount[status]++
}

func (lr *LogReport) addToQuantileEstimator(bodyBytesSent int) {
	lr.QuantileEstimator.Add(float64(bodyBytesSent))
}

func (lr *LogReport) AverageResponseSize() float64 {
	if lr.NumberRequests == 0 {
		return 0
	}

	return float64(lr.TotalResponseSize) / float64(lr.NumberRequests)
}

func (lr *LogReport) Percentile95() float64 {
	if lr.NumberRequests == 0 {
		return 0
	}

	return lr.QuantileEstimator.Get(0.95)
}
