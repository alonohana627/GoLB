package main

import (
	"GoLB/algorithms"
	"GoLB/loadbalancer"
	"GoLB/proxy"
	"net/http"
	"net/http/httputil"
	url2 "net/url"
)

func main() {
	u1, _ := url2.Parse("http://localhost:9999")
	u2, _ := url2.Parse("http://localhost:10000")

	rp1 := httputil.NewSingleHostReverseProxy(u1)
	rp2 := httputil.NewSingleHostReverseProxy(u2)

	backends := []*proxy.Backend{
		proxy.NewBackend(u1, rp1),
		proxy.NewBackend(u2, rp2),
	}

	l := loadbalancer.LoadBalancer{
		Iterator: algorithms.NewRoundRobin(backends),
	}

	http.ListenAndServe(":8080", &l)
}
