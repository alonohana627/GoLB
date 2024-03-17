package healthcheck

import (
	"net"
	"net/url"
	"sync"
	"time"
)

// TODO: remove after x health checks

type HealthChecker interface {
	PeriodicCheck()
	Check() bool
	IsAvailable() bool
	Stop()
}

type ProxyHealth struct {
	mu          sync.Mutex
	isAvailable bool
	interval    time.Duration
	addr        *url.URL
	stop        chan bool
}

func NewProxyHealth(addr *url.URL) *ProxyHealth {
	stop := make(chan bool)
	return &ProxyHealth{
		isAvailable: true,
		interval:    time.Second,
		addr:        addr,
		stop:        stop,
	}
}

func (p *ProxyHealth) PeriodicCheck() {
	checkHealth := func() {
		p.mu.Lock()
		defer p.mu.Unlock()
		p.isAvailable = p.Check()
	}

	go func() {
		t := time.NewTicker(p.interval)
		for {
			select {
			case <-t.C:
				checkHealth()
			case <-p.stop:
				return
			}
		}
	}()
}

func (p *ProxyHealth) Check() bool {
	conn, err := net.DialTimeout("tcp", p.addr.Host, time.Second)
	if err != nil {
		return false
	}

	_ = conn.Close()
	return true
}

func (p *ProxyHealth) IsAvailable() bool {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.isAvailable
}

func (p *ProxyHealth) Stop() {
	p.stop <- true
	p.isAvailable = false
	close(p.stop)
}
