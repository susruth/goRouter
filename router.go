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
	id int			// Message ID
	status string	// Message status string
	source int		// Message source router ID
	destination int // Message destination router ID
	info string		// Message content
	hops int		// No.of hops
}

// This is the type of the Router object

type Router struct {
	myId int					// Router ID
	routingTable RoutingTable	// Routing Table
	messageStatus map[int]bool	// Message delivery status map
}

// This is the type of each row in the Routing table

type RoutingEntry struct {
	connected bool		// Connection status
	destination int		// Destination router ID
	neighbour int		// Neighbour nearest to the destination
	hops int			// No.of hops to destination
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
	routers.instantiate assigns ids and empty message status maps to all the routers.
 */

func (r *Routers) instantiate() {
	for i:= 0; i < ROUTER_COUNT; i++{
		routers[i].messageStatus = make(map[int]bool)
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

	Implemented: It adds the source router ID and sends the message to the neighbour that
	is nearest to the destination router.

	Wanted to implement: Iteratively send message to all the neighbours that are connected
	to the destination router. This would be useful for the router dropping part.

	@params message Takes in the message to be broadcasted.

*/

func (r *Router) broadcast(m message) {
	channel := make(chan message)
	msg := message {
		id:m.id,
		destination:m.destination,
		source: r.myId,
		info: m.info,
		hops: m.hops,
		status:m.status,
	}
	go routers[r.routingTable[m.destination].neighbour].send(msg,channel)
	go routers[r.routingTable[m.destination].neighbour].receive(channel)
}

/*
	This go routine takes in a message and a message channel, and writes the message to the channel.

	@params message				This is the message.
	@params message channel		This is the message channel used by the goroutines

*/

func (r *Router) send(m message, c chan message) {

	m.hops += 1
	c <- m
}

/*
	listen function waits till the message to the corresponding messageID is reached, if it does it
	stops the infinite loop and returns.

	@params int					This is the message ID.

*/

func (r *Router) listen(id int) {
	for (true) {
		if (routers[1].messageStatus[id]){
			fmt.Println("Success")
			return
		}
	}
}

/*
	This go routine takes in a message channel, reads the output if it is  not for itself it
	rebroadcasts it, otherwise it will check whether it's status is "Delivered" If it is then
	it set's the message delivery status to true. Otherwise it sends a delivered status to the
	sender.

	@params message channel		This is the message channel used by the goroutines

*/

func (r *Router) receive(c chan message) {
	msg := <- c
	if (r.myId == msg.destination){
		if (msg.status == "Delivered"){
			r.messageStatus[msg.id] = true
		}else{
			success := message{
				id: msg.id,
				destination:msg.source,
				hops: 0,
				info: msg.info,
				status:"Delivered",
			}
			r.broadcast(success)
		}
	}else{
		r.broadcast(msg)
	}
}