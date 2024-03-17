package loadbalancer

import (
	"GoLB/algorithms"
	"net/http"
)

type LoadBalancer struct {
	Iterator algorithms.Iterator
}

func (l *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b := l.Iterator.GetNextBackend()
	b.ServeHTTP(w, r)
}

func NewLoadBalancer(iter algorithms.Iterator) *LoadBalancer {
	return &LoadBalancer{Iterator: iter}
}
