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
	go node3()
	go node4()
	for {
	}
}

func node0() {
	mySelf := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000")
	rt := NewRoutingTable(mySelf)

	node1 := NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8001")

	rt.AddContact(mySelf)
	rt.AddContact(node1)

	file := NewFile("/data/1.txt", true, []byte("hola"))

	channel := make(chan []Contact)
	kademlia := NewKademlia(*rt, 20, 3, channel)
	msgID := RandomInt()

	go kademlia.network.Listen("localhost", 8000)
	kademlia.Store(file)

	time.Sleep(5000000000)
	kademlia.LookupData(Hash(file.Path), msgID)
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

	//	data := []byte("Data")
	//	msgID := RandomInt()

	//time.Sleep(5000000000)
	//fmt.Println("-------------------------------LookUpDataInDHT------------------------------------------")

	//kademlia.LookupData(Hash(data), msgID)
}

func node2() {
	mySelf := NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002")
	rt := NewRoutingTable(mySelf)
	node3 := NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8003")
	//node4 := NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8004")
	rt.AddContact(mySelf)
	rt.AddContact(node3)
	//rt.AddContact(node4)

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

	//data := []byte("Data")
	channel := make(chan []Contact)
	kademlia := NewKademlia(*rt, 20, 3, channel)

	go kademlia.network.Listen("localhost", 8004)
	//kademlia.Store(data)

	/*
		fmt.Println("-------------------------------LookUpDataNoDHT------------------------------------------")

		kademlia.LookupData(Hash(data), msgID)

}
*/
