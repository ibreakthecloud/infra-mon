package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func addPingRoutes(r *gin.Engine) {
	ping := r.Group("ping")

	ping.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}

func addMetricsRoute(r *gin.Engine) {
	metrics := r.Group("/metrics")

	metrics.POST("/", func(c *gin.Context) {
		c.String(http.StatusOK, "add metrics handler")
	})
}

func addReportRoute(r *gin.Engine) {
	report := r.Group("/report")

	report.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "add report handler")
	})
}