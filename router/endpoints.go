package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Simple ping-pong router to check sanity of server
func addPingRoutes(r *gin.Engine) {
	ping := r.Group("ping")

	ping.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}

// addMetricsRoute that supports to ingest logs
func addMetricsRoute(r *gin.Engine) {
	metrics := r.Group("/metrics")

	metrics.POST("/", metricsController.StoreMetrics)
}

// addReportRoute that generates report with stats like
// highest CPU and highest memory usage
func addReportRoute(r *gin.Engine) {
	report := r.Group("/report")

	report.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "add report handler")
	})
}