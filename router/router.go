package router

import (
	"github.com/gin-gonic/gin"

)
// Endpoints constants
const(
	Metrics = "metrics"
	Report = "report"
)

func New() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	registerAllEndpoints(router)

	return router
}

func registerAllEndpoints(r *gin.Engine) {

	// Simple ping-pong router to check sanity of server
	addPingRoutes(r)
	// Adds metrics route that supports to ingest logs
	addMetricsRoute(r)
	// Adds report route that generates report with stats like
	// highest CPU and highest memory usage
	addReportRoute(r)
}
