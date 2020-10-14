package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ibreakthecloud/infra-mon/controller/metrics"
)
// Endpoints constants
const(
	Metrics = "metrics"
	Report = "report"
)

var(
	metricsController = metrics.New()
)

func New() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	registerAllEndpoints(router)

	return router
}

// registerAllEndpoints registers all of the endpoints supported by the server
func registerAllEndpoints(r *gin.Engine) {

	addPingRoutes(r)
	addMetricsRoute(r)
	addReportRoute(r)
}
