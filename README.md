# GoLB - Go Load Balancer

`The load balancer works to steer the traffic to a pool of available servers through various load balancing algorithms. If more resources are needed, additional servers can be added. Load balancers health check the application on the server to determine its availability.`

## How to Use

You can use directly:
```go
func main(){
    u1, _ := url2.Parse("http://localhost:9999")
    u2, _ := url2.Parse("http://localhost:10000")
    
    rp1 := httputil.NewSingleHostReverseProxy(u1)
    rp2	 := httputil.NewSingleHostReverseProxy(u2)
    
    backends := []*proxy.Backend{
    proxy.NewBackend(u1, rp1),
    proxy.NewBackend(u2, rp2),
    }
    
    l := loadbalancer.LoadBalancer{
    Iterator: algorithms.NewRoundRobin(backends),
    }
    
    http.ListenAndServe(":8080", &l)
}
```

Or use a config that similar to `config_example.json`:

```go
func main() {
	itConf, err := config.Parse("config_example.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = http.ListenAndServe(":9999", itConf.ToLoadBalancer())
	if err != nil {
		fmt.Println(err)
		return
	}
}
```

Currently, supports only Round Robin. Might add different algorithms in the future.