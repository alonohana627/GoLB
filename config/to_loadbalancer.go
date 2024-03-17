package config

import (
	"GoLB/algorithms"
	"GoLB/loadbalancer"
	"GoLB/proxy"
	"net/http/httputil"
	"net/url"
)

func (lc *IteratorConfig) GenerateBackends() []*proxy.Backend {
	backends := make([]*proxy.Backend, len(lc.Ips))

	for i, ip := range lc.Ips {
		urlFromIp, err := url.Parse(ip)
		if err != nil {
			// TODO: add to logger
			continue
		}

		rp := httputil.NewSingleHostReverseProxy(urlFromIp)
		backend := proxy.NewBackend(urlFromIp, rp)
		backends[i] = backend
	}

	return backends
}

func (lc *IteratorConfig) ToIterator() algorithms.Iterator {
	switch lc.AlgorithmType {
	case algorithms.RoundRobinAlgo:
		return algorithms.NewRoundRobin(lc.GenerateBackends())
	}
	return nil
}

func (lc *IteratorConfig) ToLoadBalancer() *loadbalancer.LoadBalancer {
	return loadbalancer.NewLoadBalancer(lc.ToIterator())
}
