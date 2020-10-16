package in_memory

import "github.com/ibreakthecloud/infra-mon/pkg/store"

func New() *InMemory {
	return &InMemory{
		Metrics: make(map[string][]Metrics),
	}
}

// Add methods appends the new metrics in the map against it's key(IP)
func (i *InMemory) Create(ip string, cpuUsed int, memUsed int) {
	m := Metrics{
		CPUPercentage:    cpuUsed,
		MemoryPercentage: memUsed,
	}

	i.Metrics[ip] = append(i.Metrics[ip], m)
}

// Get methods generates the report with highest CPU and memory usage
// ever recorded for the clients
func (i *InMemory) Get() []store.MaxReport {
	return generateClientMaximumReport(i)
}

// generateClientMaximumReport generates the report with clients with their
// highest CPU and memory.
// It iterates over the in-memory store and copies the highest CPU and memory
// into a store.MaxReport struct and returns the struct
func generateClientMaximumReport(i *InMemory) []store.MaxReport {
	var outputMetrics = &store.MaxReport{
		IP:        "",
		MaxCPU:    0,
		MaxMemory: 0,
	}
	var max []store.MaxReport

	// iterate over the in memory store
	for ip, metrics := range i.Metrics {
		// iterate over every metrics of the client
		for _, m := range metrics {
			if m.MemoryPercentage > outputMetrics.MaxMemory {
				outputMetrics.MaxMemory = m.MemoryPercentage
			}
			if m.CPUPercentage > outputMetrics.MaxCPU {
				outputMetrics.MaxCPU = m.CPUPercentage
			}
		}
		outputMetrics.IP = ip

		max = append(max, *outputMetrics)
		// make it empty
		outputMetrics = &store.MaxReport{}
	}
	return max
}
