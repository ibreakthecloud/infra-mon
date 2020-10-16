package store

var NewStore Store

type MaxReport struct {
	IP        string `json:"ip"`
	MaxCPU    int    `json:"max_cpu"`
	MaxMemory int    `json:"max_memory"`
}

type Store interface {
	Create(ip string, cpuUsed int, memUsed int)
	Get() []MaxReport
}
