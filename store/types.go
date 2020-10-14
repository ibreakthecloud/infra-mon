package store

// Store to store metrics
type Store struct {
	Metrics Stats
}

type Stats map[string][]Metrics

type Metrics struct {
	CPUPercentage    int
	MemoryPercentage int
}
