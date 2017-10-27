package main

import (
	"fmt"
	"os"
)

func main() {
	m := message{
		info:"Hello",
		destination:4,
		hops:0,
	}
	routers.configure("ring")
	os.Exit(0)
	// fmt.Println(routers)
	// routers[3].kill()
	routers[2].broadcast(m)
	var input string
	fmt.Scanln(&input)
}






//
//fmt.Println("-----------------------Configuring Routers In Ring Topology------------------------")
//routers.configure("ring")
//sum := 0
//min := INFINITY
//for i := 0; i < ROUTER_COUNT; i++ {
//for j := 0; j < ROUTER_COUNT; j++{
//sum += routers[i].routingTable[j].hops
//if (min > routers[i].routingTable[j].hops){
//min = routers[i].routingTable[j].hops
//}
//}
//}
//fmt.Println("Average Hops: ", sum/(ROUTER_COUNT*ROUTER_COUNT))
//fmt.Println("Minimum Hops: ", min)
//fmt.Println("-----------------------DONE---------------------------------------")
//
//fmt.Println("-----------------------Configuring Routers In Star Topology------------------------")
//routers.configure("star")
//sum = 0
//min = INFINITY
//for i := 0; i < ROUTER_COUNT; i++ {
//for j := 0; j < ROUTER_COUNT; j++{
//sum += routers[i].routingTable[j].hops
//if (min > routers[i].routingTable[j].hops){
//min = routers[i].routingTable[j].hops
//}
//}
//}
//fmt.Println("Average Hops: ", float64(sum)/float64(ROUTER_COUNT*ROUTER_COUNT))
//fmt.Println("Minimum Hops: ", min)
//fmt.Println("-----------------------DONE---------------------------------------")
//
