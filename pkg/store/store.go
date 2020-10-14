package store

var DataStore *Store

func New() *Store {
	return &Store{
		Metrics: make(map[string][]Metrics),
	}
}

// Add methods appends the new metrics in the map against it's key(IP)
func (s *Store) Add(ip string, cpuUsed int, memUsed int) {
	m := Metrics{
		CPUPercentage:    cpuUsed,
		MemoryPercentage: memUsed,
	}

	DataStore.Metrics[ip] = append(DataStore.Metrics[ip], m)
}

func (s *Store) GetUniqueMaximum() []MaxReport {
	var outputMetrics = &MaxReport{
		IP:        "",
		MaxCPU:    0,
		MaxMemory: 0,
	}
	var max []MaxReport

	for ip, metrics := range s.Metrics {
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
		outputMetrics = &MaxReport{}
	}
	return max
}