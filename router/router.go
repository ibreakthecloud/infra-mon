package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ibreakthecloud/infra-mon/controller/metrics"
	"github.com/ibreakthecloud/infra-mon/controller/report"
)

// Endpoints constants
const(
	Metrics = "metrics"
	Report = "report"
)

var(
	metricsController = metrics.New()
	reportController = report.New()
)

// New instantiates a new gin router to handle API requests
func New() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	registerAllEndpoints(router)

	return router
}

// registerAllEndpoints registers all of the endpoints supported by the server
func registerAllEndpoints(r *gin.Engine) {

	// registers the below endpoints
	addPingRoutes(r)
	addMetricsRoute(r)
	addReportRoute(r)
}
