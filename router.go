/*
Author : Nadimpalli Susruth
UID: U6106064
Email : susruth.praker@gmail.com
Created Date : 14/10/17
Updated Date : 27/10/17
*/

package main

import (
	"fmt"
)


/*

	My implementation does not support dropping routers. It only saves the neighbour
	that is closest to the destination in the routing table.

 */



// The number of routers can be set over here.
const ROUTER_COUNT = 25

// A very large number simulating infinite.
const INFINITY = 99999999

var routers Routers


// This is the type of the message being passed.

type message struct {
	destination int
	info string
	hops int
}

// This is the type of the Router object

type Router struct {
	myId int
	routingTable RoutingTable
}

// This is the type of each row in the Routing table

type RoutingEntry struct {
	connected bool
	destination int
	neighbour int
	hops int
}

// This is the type of the routing table.

type RoutingTable [ROUTER_COUNT]RoutingEntry

// This is the array of all the routers

type Routers [ROUTER_COUNT]Router



//////// Router Discovery Phase.

/*
	routers.configure takes in a string and initializes the routers with the given topology.
	@params string
 */
func (r *Routers) configure(topology string) {
	r.instantiate()
	r.setup(topology)
	r.connect()
}

/*
	routers.instantiate assigns ids to all the routers.
 */

func (r *Routers) instantiate() {
	for i:= 0; i < ROUTER_COUNT; i++{
		r[i].myId = i
	}
}

/*
	routers.setup assigns a basic routing tables depending on the given topology to all the routers.
	this routing table only consists of the info about their immediate neighbours.
	@params string
 */

func (r *Routers) setup(topology string){
	for i:= 0; i < ROUTER_COUNT; i++{
		r[i].routingTable = configureConnections(topology, r[i].myId)
	}
}

/*
	routers.connect iteratively completes the routing tables depending
	on their neighbours routing table entries.

 */

func (r *Routers) connect(){
	for a:= 0; a < ROUTER_COUNT; a++ {
		for i := 0; i < ROUTER_COUNT; i++ {
			for j := 0; j < ROUTER_COUNT; j++ {
				if (r[i].routingTable[j].hops == 1){
					for k := 0; k < ROUTER_COUNT; k++ {
						if r[i].routingTable[k].hops > r[j].routingTable[k].hops+1 {
							r[i].routingTable[k].neighbour = r[j].myId
							r[i].routingTable[k].hops = r[j].routingTable[k].hops + 1
							r[i].routingTable[k].connected = true
						}

					}
				}
			}
		}
	}
}


/////////// Router Message Passing Part


/*

	Implemented: It sends the message to the neighbour that is nearest to the destination router.

	Wanted to implement: Iteratively send message to all the neighbours that are connected to the destination router.
	This would be useful for the router dropping part.

	@params message Takes in the message to be broadcasted.

*/

func (r Router) broadcast(m message) {
	channel := make(chan message)
	go routers[r.routingTable[m.destination].neighbour].send(m,channel)
	go routers[r.routingTable[m.destination].neighbour].receive(channel)
}

/*
	This go routine takes in a message and a message channel, and writes the message to the channel.

	@params message				This is the message.
	@params message channel		This is the message channel used by the goroutines

*/

func (r Router) send(m message, c chan message) {
	m.hops += 1
	c <- m
}

/*
	This go routine takes in a message channel, reads the outpur if it is for itself, it exits
	otherwise it rebrodcasts it.

	@params message channel		This is the message channel used by the goroutines

*/

func (r Router) receive(c chan message) {
	msg := <- c
	if (r.myId == msg.destination){
		fmt.Println("Message: ", msg.info,"\nHops: ",msg.hops)
	}else{
		r.broadcast(msg)
	}
}