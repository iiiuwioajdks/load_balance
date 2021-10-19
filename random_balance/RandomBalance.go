package main

/*
随机法
*/

import (
	"fmt"
	"math/rand"
	"time"
)

type RandomBalance struct {
	Host string
}

func (r *RandomBalance) DoBalance(rbs []*RandomBalance) (rb *RandomBalance) {
	index := rand.Intn(len(rbs))
	return rbs[index]
}

func main() {
	rServers := []*RandomBalance{
		{"121.40.193.222:9898"},
		{"121.40.193.223:9785"},
		{"121.40.193.229:8698"},
	}
	rServer := new(RandomBalance)
	for {
		rs := rServer.DoBalance(rServers)
		fmt.Println(rs.Host)
		time.Sleep(time.Second)
	}
}
