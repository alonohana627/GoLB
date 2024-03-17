package main

import (
	"GoLB/config"
	"fmt"
	"net/http"
)

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
