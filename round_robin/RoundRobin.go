package main

/*
轮询法
*/

import (
	"fmt"
	"time"
)

type RoundRobin struct {
	servers []string
	now     int
}

func (r *RoundRobin) DoBalance() string {
	now := (r.now + 1) % len(r.servers)
	r.now = now
	return r.servers[r.now]
}

func main() {
	r := new(RoundRobin)
	r.servers = []string{"121.40.193.222:9898", "121.40.193.223:9785", "121.40.193.229:8698"}
	r.now = -1

	for {
		fmt.Println(r.DoBalance())
		time.Sleep(time.Second)
	}
}
