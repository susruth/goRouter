/*
Author : Nadimpalli Susruth
UID: U6106064
Email : susruth.praker@gmail.com
Created Date : 14/10/17
Updated Date : 27/10/17
*/

package main

/*
	configureConnections is a switch statement that can be used to select the required topology.
	And it returns the routing table of the given router ID as the output.

	@params string takes in the topology string and the router's id.
	@returns It returns the routing table of the given router.

 */

func configureConnections(topology string, id int) RoutingTable {
	var table = RoutingTable{}
	switch (topology){
	case "ring":
		table = ring(id)
		break
	case "line":
		table = line(id)
		break
	case "star":
		table = star(id)
		break
	case "fullyConnected":
		table = fullyConnected(id)
		break
	}
	return table
}

/*
	Ring Topology:
		In this topology all the routers are connected in a chain.

	ring takes in the id of the router and returns the routing table for the Ring Topology.
	@params id : int takes in the id of  the router and returns a routing table.
	@returns RoutingTable routing table of the given router.

 */


func ring(id int) RoutingTable {
	table := RoutingTable{}
	for j:= 0; j < ROUTER_COUNT; j++ {
		table[j].destination = j
		table[j].hops = INFINITY
		if ((ROUTER_COUNT + id - j) % ROUTER_COUNT == 1 || (ROUTER_COUNT + id - j) % ROUTER_COUNT == ROUTER_COUNT - 1) {
			table[j].connected = true
			table[j].neighbour = j
			table[j].hops = 1
		}else if (j == id){
			table[j].connected = true
			table[j].hops = 0
			table[j].neighbour = j
		}
	}
	return table
}

/*
	Line/Bus Topology:
		In this topology all the routers are connected in a line.

	line takes in the id of the router and returns the routing table for the Line/Bus Topology.
	@params id : int takes in the id of  the router and returns a routing table.
	@returns RoutingTable routing table of the given router.

*/

func line(id int) RoutingTable {
	table := RoutingTable{}
	for j:= 0; j < ROUTER_COUNT; j++ {
		table[j].destination = j
		table[j].hops = INFINITY
		if (id + 1 == j || id - 1 == j) {
			table[j].connected = true
			table[j].neighbour = j
			table[j].hops = 1
		}else if (j == id){
			table[j].connected = true
			table[j].hops = 0
			table[j].neighbour = j
		}
	}
	return table
}

/*
	Star Topology:
		In this topology all the routers are connected to a hub router which relays information
	between any two routers on the network.

	star takes in the id of the router and returns the routing table for the Star Topology.
	@params id : int takes in the id of  the router and returns a routing table.
	@returns RoutingTable routing table of the given router.

*/

func star(id int) RoutingTable {
	table := RoutingTable{}


	if id == 0{
		for j:= 0; j < ROUTER_COUNT; j++ {
			table[j].destination = j;
			table[j].hops = 1;
			table[j].connected = true;
		}
		table[id].hops = 0;
	}else {
		for j:= 0; j < ROUTER_COUNT; j++ {
			table[j].destination = j;
			table[j].hops = INFINITY;
			table[j].connected = false;
			if (j == 0){
				table[j].hops = 1;
				table[j].connected = true;
			}
		}
		table[id].hops = 0;
	}
	return table
}

/*
	FullyConnected Topology:
		In this topology all the routers are connected to each other.

	fullyConnected takes in the id of the router and returns the routing table for the FullyConnected Topology.
	@params id : int takes in the id of  the router and returns a routing table.
	@returns RoutingTable routing table of the given router.

*/


func fullyConnected(id int) RoutingTable {
	table := RoutingTable{}
	for j:= 0; j < ROUTER_COUNT; j++ {
		table[j].destination = j;
		table[j].hops = 1;
		table[j].connected = true;
		if (id == j){
			table[id].hops = 0;
		}
	}
	return table
}