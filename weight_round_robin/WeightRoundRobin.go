package main

/*
加权轮询法
*/
import (
	"fmt"
	"time"
)

type WeightServer struct {
	Host          string
	CurrentWeight int // 当前的权重
	Weight        int // 权重
}

func (w *WeightServer) DoBalance(wServers []*WeightServer) (ws *WeightServer) {
	allWeight := 0 // 总权重
	if len(wServers) == 0 {
		return nil
	}

	for _, wServer := range wServers {
		allWeight += wServer.Weight
		wServer.CurrentWeight += wServer.Weight

		if ws == nil || ws.CurrentWeight < wServer.CurrentWeight {
			ws = wServer
		}
	}

	// 实现平滑的的加权轮询,使各个服务的CurrentWeight处在一个T为总权值的周期循环
	// 因为每个都减掉总权重，又慢慢的加上自己的权重，所有会保持循环
	// 即权重为3，2，1的话，出现出现的序列为 ABACBA
	ws.CurrentWeight -= allWeight
	return
}

func main() {
	wServers := []*WeightServer{
		{"121.40.193.222:9898", 0, 2},
		{"121.40.193.223:9785", 0, 1},
		{"121.40.193.229:8698", 0, 1},
	}
	w := new(WeightServer)
	for {
		wServer := w.DoBalance(wServers)

		fmt.Println(wServer.Host+" ", wServer.CurrentWeight)
		time.Sleep(time.Second)
	}
}
