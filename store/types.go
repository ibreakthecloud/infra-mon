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

type MaxReport struct {
	IP        string `json:"ip"`
	MaxCPU    int    `json:"max_cpu"`
	MaxMemory int    `json:"max_memory"`
}
