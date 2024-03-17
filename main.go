package main

import (
	"GoLB/config"
	"fmt"
	"net/http"
)

func main() {
	//u1, _ := url2.Parse("http://localhost:9999")
	//u2, _ := url2.Parse("http://localhost:10000")
	//
	//rp1 := httputil.NewSingleHostReverseProxy(u1)
	//rp2 := httputil.NewSingleHostReverseProxy(u2)
	//
	//backends := []*proxy.Backend{
	//	proxy.NewBackend(u1, rp1),
	//	proxy.NewBackend(u2, rp2),
	//}
	//
	//l := loadbalancer.LoadBalancer{
	//	Iterator: algorithms.NewRoundRobin(backends),
	//}
	//
	//http.ListenAndServe(":8080", &l)

	itConf, err := config.Parse("config_example.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	http.ListenAndServe(":9999", itConf.ToLoadBalancer())
}
