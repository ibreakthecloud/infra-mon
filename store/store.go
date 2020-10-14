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