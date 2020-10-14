package metrics

type Metrics struct {
	CPUPercentage int `json:"percentage_cpu_used"`
	MemoryPercentage int `json:"percentage_memory_used"`
}