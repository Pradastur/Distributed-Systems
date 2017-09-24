package main

import (
	"fmt"
	"testing"
	"time"
)

func TestRoutingTable(t *testing.T) {
	go node0()
	go node1()
	go node2()
	for {
	}
<<<<<<< HEAD

	kademlia:=NewKademlia(*rt,2,1)
	go kademlia.ServerThread("8000")
<<<<<<< HEAD
=======
	fmt.Println(SendPingMessage(srcContact, srcContact))
>>>>>>> b799241e9953bb156fdedc5e994fabaecb300906
=======
}

func node0() {
	mySelf := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000")
	rt := NewRoutingTable(mySelf)

	node1 := NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8001")

	rt.AddContact(mySelf)
	rt.AddContact(node1)
	channel := make(chan []Contact)
	kademlia := NewKademlia(*rt, 20, 3, channel)
	data := []byte("Some rdm data")
	go kademlia.network.Listen("localhost", 8000)
	kademlia.Store(data)
}

func node1() {
	mySelf := NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8001")
	node2 := NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8002")

	rt := NewRoutingTable(mySelf)

	rt.AddContact(mySelf)
	rt.AddContact(node2)

	channel := make(chan []Contact)
	kademlia := NewKademlia(*rt, 20, 3, channel)

	go kademlia.network.Listen("localhost", 8001)
}

func node2() {
	mySelf := NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8002")
	node1 := NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8001")

	rt := NewRoutingTable(mySelf)
	rt.AddContact(mySelf)
	rt.AddContact(node1)
	channel := make(chan []Contact)
	kademlia := NewKademlia(*rt, 20, 3, channel)
	go kademlia.network.Listen("localhost", 8002)
	time.Sleep(5000000000)
	fmt.Println("---------------------------SECOND PART---------------------------")
	data := []byte("Some rdm data")
	msgID := RandomInt()
	kademlia.LookupData(Hash(data), msgID)
>>>>>>> 767f8db8bd364d4c57f29672aff47286b89d27ee
}
