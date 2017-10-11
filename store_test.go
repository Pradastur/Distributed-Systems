package main

/*
import (
	"testing"
)

func Test1(t *testing.T) {
	go node0()
	go node1()
	go node2()
	go node3()
	go node4()
	for {
	}
}

func node0() {
	mySelf := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000")
	rt := NewRoutingTable(mySelf)

	node1 := NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8001")
	node2 := NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002")
	node3 := NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8003")

	rt.AddContact(mySelf)
	rt.AddContact(node1)
	rt.AddContact(node2)
	rt.AddContact(node3)

	channel := make(chan []Contact)
	kademlia := NewKademlia(*rt, 20, 3, channel)
	go kademlia.network.Listen("localhost", 8000)

	//1. (LookupContact)Looking for node1 (already have it in my routing table)
	//kademlia.LookupContact(&node1, 1234)
	//2. (LookupContact)Looking for node4 (don't have in my routing table)
	//kademlia.LookupContact(&node4, 1234)
	//3. (LookupData)Looking for data in node0 (already have it)
	//kademlia.LookupData(Hash(data), 1234)
	//4. (LookupData)Looking for data in node1 (don't have it but already
	//have the destiny node in my routing table)
	//kademlia.LookupData(Hash(data1), msgID)
	//5. (LookupData)Looking for data in node4 (don't have it and don't
	//have the destiny node in my routing table)
	//kademlia.LookupData(Hash(data4), 1234)
}

func node1() {
	mySelf := NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8001")
	node0 := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000")

	rt := NewRoutingTable(mySelf)

	rt.AddContact(mySelf)
	rt.AddContact(node0)

	channel := make(chan []Contact)
	file := NewFile("/data/1.txt", true, []byte("hola"))
	kademlia := NewKademlia(*rt, 20, 3, channel)
	go kademlia.network.Listen("localhost", 8001)
	kademlia.Store(file)

}

func node2() {
	mySelf := NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002")
	rt := NewRoutingTable(mySelf)
	node4 := NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8004")

	rt.AddContact(mySelf)
	rt.AddContact(node4)
	channel := make(chan []Contact)
	kademlia := NewKademlia(*rt, 20, 3, channel)
	go kademlia.network.Listen("localhost", 8002)
}

func node3() {
	mySelf := NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8003")
	rt := NewRoutingTable(mySelf)

	rt.AddContact(mySelf)
	channel := make(chan []Contact)
	kademlia := NewKademlia(*rt, 20, 3, channel)
	go kademlia.network.Listen("localhost", 8003)
}

func node4() {
	mySelf := NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8004")

	rt := NewRoutingTable(mySelf)

	rt.AddContact(mySelf)
	channel := make(chan []Contact)
	kademlia := NewKademlia(*rt, 20, 3, channel)
	go kademlia.network.Listen("localhost", 8004)
}
*/
