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

	// TODO: add missing implementation

	return router
}
