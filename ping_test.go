package main

/*
import (
	"testing"
	"time"
)

func Test1(t *testing.T) {
	go node0()
	go node1()
	go node2()
	for {
	}
}

func node0() {
	mySelf := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000")
	rt := NewRoutingTable(mySelf)

	node1 := NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8001")
	node2 := NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002")

	rt.AddContact(mySelf)

	channel := make(chan []Contact)
	kademlia := NewKademlia(*rt, 20, 3, channel)

	go kademlia.network.Listen("localhost", 8000)

	kademlia.network.SendPingMessage(&node1)

	time.Sleep(5000000000)
	kademlia.LookupContact(&node2, 1234)

}

func node1() {
	mySelf := NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8001")
	node2 := NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002")

	rt := NewRoutingTable(mySelf)

	rt.AddContact(mySelf)
	rt.AddContact(node2)

	channel := make(chan []Contact)
	kademlia := NewKademlia(*rt, 20, 3, channel)
	go kademlia.network.Listen("localhost", 8001)
}

func node2() {
	mySelf := NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002")

	rt := NewRoutingTable(mySelf)

	rt.AddContact(mySelf)

	channel := make(chan []Contact)
	kademlia := NewKademlia(*rt, 20, 3, channel)
	go kademlia.network.Listen("localhost", 8002)
}*/
