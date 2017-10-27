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
	"time"
	"math/rand"
)

/*
	I implemented a synchronous message passing system, where the destination router sends back
	a confirmation request when it gets the message. The source router listens for the confirmation
	by using the message id.

	The packages I used are "fmt" to print lines, "math/rand" package to create random numbers
	for the message IDs and "time" to print the time taken to configure.

	Instructions:
		1. run "go run *" from the main folder to run this code.

 */

func main() {

	// Message ID for synchronous message passing.
	messageId := rand.Intn(100)

	// New message creation.
	m := message{
		id:messageId,
		info:"Hello",
		destination:2,
		hops:0,
	}


	// Stats of all the topologies.
	stats("ring")
	stats("line")
	stats("star")
	stats("fullyConnected")

	// Message passing example.
	routers.configure("ring")
	routers[1].broadcast(m)

	// Waits for the delivery success of the message
	routers[1].listen(messageId)
}

/*

	stats function takes in topology string, and configures routers and prints the metrics
	mainly "Average" and "Max" hops.
	@params string takes in the topology string.

 */


func stats(topology string){
	fmt.Println()
	fmt.Println("-----------------------Configuring routers in",topology,"topology starts------------------------")
	start := time.Now()
	routers.configure(topology)
	t := time.Now()
	fmt.Println("-----------------------Configuring routers in",topology,"topology ended in ",t.Sub(start),"---------")

	sum := 0
	max := 0
	for i := 0; i < ROUTER_COUNT; i++ {
		for j := 0; j < ROUTER_COUNT; j++{
			if (i != j) {
				sum += routers[i].routingTable[j].hops
				if (max < routers[i].routingTable[j].hops) {
					max = routers[i].routingTable[j].hops
				}
			}
		}
	}
	fmt.Println()
	fmt.Println("Average", float64(sum)/float64(ROUTER_COUNT*(ROUTER_COUNT-1)))
	fmt.Println("Max", max)
	fmt.Println()
	fmt.Println("-----------------------------------------DONE-------------------------------------------------------")
	fmt.Println()
}