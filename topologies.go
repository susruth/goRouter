package main

import "fmt"

func configureConnections(topology string, id int) [ROUTER_COUNT]RoutingEntry {
	var table = [ROUTER_COUNT]RoutingEntry{}
	switch (topology){
	case "ring":
		table = ring(id)
		break
	//case "line":
	//	line()
	//	break
	//case "star":
	//	table = star(id)
	//	break
	////case "fullyConnected":
	//	fullyConnected()
	//	break
	//case "tree":
	//	tree()
	//	break
	}
	return table
}
//func ring(rtable chan [ROUTER_COUNT]RoutingEntry) {
//	var routingTable :[ROUTER_COUNT]chan RoutingEntry);
//	for i := 1 ; i < ROUTER_COUNT; i++ {
//		for j := 1; j < ROUTER_COUNT; j++ {
//			routingTable[i].destination = j
//			if ((ROUTER_COUNT + i - j) % ROUTER_COUNT == 1 || (ROUTER_COUNT + i - j) % ROUTER_COUNT == ROUTER_COUNT - 1) {
//				routingTable[i].connected = true
//			}
//		}
//	}
//	rtable = routingTable;
//}


func ring(id int) [ROUTER_COUNT]RoutingEntry {
	table := [ROUTER_COUNT]RoutingEntry{}
	for j:= 0; j < ROUTER_COUNT; j++ {
		table[j].destination = j
		for i :=  0; i < 2; i++ {
			table[j].neighbours[i].hops = INFINITY
			if ((ROUTER_COUNT + id - j) % ROUTER_COUNT == 1 || (ROUTER_COUNT + id - j) % ROUTER_COUNT == ROUTER_COUNT - 1) {
				table[j].connected = true
				table[j].neighbours[i].id = j
				table[j].neighbours[i].hops = 1
			}else if (j == id){
				table[j].connected = true
				table[j].neighbours[i].hops = 0
				table[j].neighbours[i].id = j
			}
		}
	}
	fmt.Println(table)
	return table
}


//func star(id int) [ROUTER_COUNT]RoutingEntry {
//	table := [ROUTER_COUNT]RoutingEntry{}
//	if (id == 0){
//		for j:= 1; j < ROUTER_COUNT; j++ {
//			table[j].connected = true
//			table[j].destination = j
//			table[j].neighbour.hops = 1
//		}
//	} else {
//		for j:= 1; j < ROUTER_COUNT; j++ {
//			table[j].destination = j
//			table[j].neighbour.hops = INFINITY
//		}
//		table[0].connected = true
//		table[0].destination = 0
//		table[0].neighbour.hops = 1
//	}
//	table[id].neighbour.hops = 0
//	table[id].connected = true
//	table[id].neighbour.id = id
//	return table
//}
//
//func line() {
//	for i := 1 ; i < ROUTER_COUNT; i++{
//		routers[i].myId = i;
//		routers[i].connections = [ROUTER_COUNT]int{(ROUTER_COUNT-1+i) % ROUTER_COUNT,(ROUTER_COUNT+1+i) % ROUTER_COUNT}
//	}
//	routers[1].connections =  [ROUTER_COUNT]int{2}
//	routers[ROUTER_COUNT-1].connections =  [ROUTER_COUNT]int{ROUTER_COUNT-2}
//
//}
//
//func fullyConnected(){
//	for i := 1 ; i < ROUTER_COUNT; i++{
//		routers[i].myId = i;
//		for j := 0; j < ROUTER_COUNT; j++ {
//			routers[i].connections[j] = j
//		}
//	}
//	fmt.Println(routers)
//}
//
//
//func tree() {
//	stars := int(math.Floor(math.Sqrt(float64(ROUTER_COUNT))))
//	for i:= 1; i < ROUTER_COUNT; i += stars {
//		routers[i].myId = i
//		for j := 1; j < stars && i+j < ROUTER_COUNT ; j++ {
//			routers[i].connections[i+j] = i+j
//			routers[i+j].myId = i+j
//			routers[i+j].connections[i] = i
//		}
//		for k:= 1; k < ROUTER_COUNT; k += stars {
//			routers[i].connections[k] = k
//		}
//	}
//}
//
//func torus() {
//	cycles := int(math.Floor(math.Cbrt(float64(ROUTER_COUNT))))
//	for i:= 1; i < ROUTER_COUNT; i += cycles^2 {
//		for j:= 1; j< cycles^2 && i+j < ROUTER_COUNT; j += cycles {
//			for k:= 1; k< cycles && i+j+k < ROUTER_COUNT; k++ {
//				routers[i].connections[i+j+k] = i+j+k
//				routers[i+j+k].myId = i+j+k
//				routers[i+j+k].connections[i] = i
//			}
//		}
//	}
//
//}