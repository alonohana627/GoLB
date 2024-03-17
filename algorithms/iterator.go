package algorithms

import "GoLB/proxy"

// TODO: RemoveBackend(index int)
type Iterator interface {
	GetNextBackend() *proxy.Backend
	AddBackend(backend *proxy.Backend)
	PoolSize() int
}
