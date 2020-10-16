package in_memory

// Store to store metrics
type InMemory struct {
	Metrics Stats
}

type Stats map[string][]Metrics

type Metrics struct {
	CPUPercentage    int
	MemoryPercentage int
}
