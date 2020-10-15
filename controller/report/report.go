package report

import (
	"github.com/gin-gonic/gin"
	"github.com/ibreakthecloud/infra-mon/pkg/store"
)
// New creates new report object
func New() *Report{
	return &Report{}
}

func (r *Report) GenerateReport (c *gin.Context) {
	report := store.NewStore.Get()
	c.JSON(200, report)
}
