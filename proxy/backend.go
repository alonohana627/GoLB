package proxy

import (
	"GoLB/healthcheck"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

type IBackend interface {
	Stop()
	IsAvailable() bool
	GetLoad() int32
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func (b *Backend) Stop() {
	b.health.Stop()
}

func (b *Backend) IsAvailable() bool {
	return b.health.IsAvailable()
}

func (b *Backend) GetLoad() int32 {
	return atomic.LoadInt32(&b.load)
}

func (b *Backend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt32(&b.load, 1)
	defer atomic.AddInt32(&b.load, -1)

	b.proxy.ServeHTTP(w, r)
}

type Backend struct {
	health *healthcheck.ProxyHealth
	proxy  *httputil.ReverseProxy
	load   int32
}

func NewBackend(addr *url.URL, rp *httputil.ReverseProxy) *Backend {
	healthCheck := healthcheck.NewProxyHealth(addr)
	healthCheck.PeriodicCheck()

	return &Backend{
		health: healthCheck,
		proxy:  rp,
		load:   0,
	}
}
