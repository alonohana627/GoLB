package algorithms

import "GoLB/proxy"

// TODO: RemoveBackend(index int)
// TODO: GenerateBackends() for general status check

type Iterator interface {
	GetNextBackend() *proxy.Backend
	AddBackend(backend *proxy.Backend)
	PoolSize() int
}
