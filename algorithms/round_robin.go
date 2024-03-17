package algorithms

import (
	"GoLB/proxy"
	"sync"
)

type RoundRobin struct {
	backends     []*proxy.Backend
	mu           sync.Mutex
	currentIndex int
}

func NewRoundRobin(backends []*proxy.Backend) *RoundRobin {
	return &RoundRobin{
		backends: backends,
	}
}

// GetNextBackend
// TODO: handle when all servers aren't working
func (r *RoundRobin) GetNextBackend() *proxy.Backend {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.currentIndex++
	r.currentIndex = r.currentIndex % len(r.backends)

	for !r.backends[r.currentIndex].IsAvailable() {
		r.currentIndex++
		r.currentIndex = r.currentIndex % len(r.backends)
	}

	return r.backends[r.currentIndex]
}

func (r *RoundRobin) AddBackend(backend *proxy.Backend) {
	r.backends = append(r.backends, backend)
}

func (r *RoundRobin) PoolSize() int {
	return len(r.backends)
}
