package stats

import "time"

type Stats struct {
	Start  time.Time
	MaxCPU float64
	MinCPU float64
	Init   bool
}

func NewStats() *Stats {
	return &Stats{
		Start: time.Now(),
	}
}
func (s *Stats) Update(cpu float64) {
	if !s.Init {
		s.MinCPU = cpu
		s.MaxCPU = cpu
		s.Init = true
		return
	}
	if cpu > s.MaxCPU {
		s.MaxCPU = cpu
	}
	if cpu < s.MinCPU {
		s.MinCPU = cpu
	}
}

func (s *Stats) Uptime() time.Duration {
	return time.Since(s.Start)
}
