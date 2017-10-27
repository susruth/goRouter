package main

import (
	"fmt"
)

const ROUTER_COUNT = 5
const INFINITY = 99999999
var routers Routers

type message struct {
	destination int
	info string
	hops int
}

type Router struct {
	myId int
	routingTable [ROUTER_COUNT]RoutingEntry
}

type RoutingEntry struct {
	connected bool
	destination int
	neighbours []Neighbour
}

type Neighbour struct {
	hops int
	id int
}
type Routers [ROUTER_COUNT]Router

//////// Discovery

func (r *Routers) configure(topology string) {
	r.instantiate()
	r.setup(topology)
	r.connect()
}

func (r *Routers) instantiate() {
	for i:= 0; i < ROUTER_COUNT; i++{
		r[i].myId = i
	}
}

func (r *Routers) setup(topology string){
	for i:= 0; i < ROUTER_COUNT; i++{
		r[i].routingTable = configureConnections(topology, r[i].myId)
	}
}


func (r *Routers) connect(){
	for i:= 0; i < ROUTER_COUNT; i++{
		for j:= 0; j < ROUTER_COUNT; j++{
			for k:= 0; k < len(r[i].routingTable[j].neighbours); k++{
				if (r[i].routingTable[j].neighbours[k].hops > minHops(r[r[i].routingTable[j].neighbours[k].id].routingTable[j].neighbours) + 1){
					r[i].routingTable[j].neighbours[k].hops = minHops(r[r[i].routingTable[j].neighbours[k].id].routingTable[j].neighbours) + 1
					fmt.Println(r[i].routingTable[j].neighbours[k])
				}
			}
		}
	}
}


//func (r *Routers) connect(){
//	for i:= 0; i < ROUTER_COUNT; i++{
//		for j:= 0; j < ROUTER_COUNT; j++{
//			for k:= 0; k < ROUTER_COUNT; k++{
//				if (r[i].routingTable[k].neighbour.hops > r[i].routingTable[j].neighbour.hops + 1){
//					r[i].routingTable[k].neighbour.id = r[i].routingTable[j].neighbour.id
//					r[i].routingTable[k].neighbour.hops = r[i].routingTable[j].neighbour.hops + 1
//					r[i].routingTable[k].connected = true
//
//				}
//
//			}
//		}
//	}
//
//}

//
//func connect(){
//	for _,router := range routers{
//		for _,entry := range router.routingTable{
//			for _,neighbour := range entry.neighbours{
//				if(routers[neighbour.id].routingTable[entry.destination].connected && neighbour.hops > minHops(routers[neighbour.id].routingTable[entry.destination].neighbours)+1){
//					neighbour.hops = minHops(routers[neighbour.id].routingTable[entry.destination].neighbours)+1
//				}
//			}
//			entry.connected = true
//		}
//	}
//}

func minHops(neighbours []Neighbour) int {
	hops := INFINITY
	for _,neighbour := range neighbours{
		if (hops > neighbour.hops){
			hops = neighbour.hops
		}
	}
	return hops
}


/////////// Message Passing


func (r Router) broadcast(m message) {
	channel := make(chan message)
	nearestNeighbour := findNearestNeighbour(r.routingTable[m.destination]);
	go routers[nearestNeighbour.id].send(m,channel)
	go routers[nearestNeighbour.id].receive(channel)
}

func (r Router) send(m message, c chan message) {
	m.hops += 1
	c <- m
}

func (r Router) receive(c chan message) {
	msg := <- c
	if (r.myId == msg.destination){
		fmt.Println("Message: ", msg.info,"\nHops: ",msg.hops)
	}else{
		r.broadcast(msg)
	}
}

func (r Router) kill(){
	for i := 0; i < ROUTER_COUNT; i++{
		for _, neighbour := range routers[i].routingTable[r.myId].neighbours{
			neighbour.hops = INFINITY
			routers[i].routingTable[r.myId].connected = false
		}
		for j := 0; j < ROUTER_COUNT; j++{
			for _, neighbour := range routers[i].routingTable[j].neighbours {
				if (neighbour.id == r.myId) {
					neighbour.hops = INFINITY
				}
			}
		}
	}
}


////////// Helper functions

func findNearestNeighbour(entry RoutingEntry) Neighbour {
	nearest := Neighbour{
		id: 0,
		hops:INFINITY,
	}
	for _,neighbour := range entry.neighbours{
		if (nearest.hops > neighbour.hops){
			nearest = neighbour
		}
	}
	return nearest
}