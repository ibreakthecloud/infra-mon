package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/ibreakthecloud/infra-mon/pkg/store"
	"net/http"
)

// New instantiates a new empty Metrics object
func New() *Metrics {
	return &Metrics{}
}

func (m *Metrics) StoreMetrics (c *gin.Context) {

	// verify if request payload is valid
	err := c.BindJSON(m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"unable to parse request payload"})
		return
	}

	// validate MemoryPercentage
	if !isPercent(m.MemoryPercentage) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid memory usage percentage",
		})
		return
	}

	// validate CPUPercentage
	if !isPercent(m.CPUPercentage) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid CPU usage percentage",
		})
		return
	}

	// add new metrics to store
	store.NewStore.Create(c.ClientIP(), m.CPUPercentage, m.MemoryPercentage)

	c.JSON(http.StatusOK, gin.H{
		"message": "Metric Ingested Successfully",
	})
	return
}

// isPercent validates if the provided integer is a valid percentage
// i.e. <between 0-100>
func isPercent(p int) bool {
	if p >= 0 && p <= 100 {
		return true
	}
	return false
}