package main

import (
	"strconv"
	"testing"
)

func Test1(t *testing.T) {

	go node0()
	for i := 1; i <= 98; i++ {
		if i < 10 {
			port := i
			KademliaIDName := NewKademliaID("000000000000000000000000000000000000000" + strconv.Itoa(port))
			portNext := i + 1
			KademliaIDNameNext := NewKademliaID("000000000000000000000000000000000000000" + strconv.Itoa(portNext))
			//fmt.Println(KademliaIDName.String())
			go nextNode(KademliaIDName, KademliaIDNameNext, port, portNext)
			//KademliaIDnode++
		} else {
			port := i
			KademliaIDName := NewKademliaID("00000000000000000000000000000000000000" + strconv.Itoa(port))
			portNext := i + 1
			KademliaIDNameNext := NewKademliaID("00000000000000000000000000000000000000" + strconv.Itoa(portNext))

			//fmt.Println(KademliaIDName)
			//KademliaIDnode++
			go nextNode(KademliaIDName, KademliaIDNameNext, port, portNext)
		}
	}
	go finalNode()
	for {
	}
}

func node0() {
	mySelf := NewContact(NewKademliaID("0000000000000000000000000000000000000000"), "localhost:8000")
	next := NewContact(NewKademliaID("0000000000000000000000000000000000000001"), "localhost:8001")

	rt := NewRoutingTable(mySelf)

	rt.AddContact(mySelf)
	rt.AddContact(next)

	channel := make(chan []Contact)
	kademlia := NewKademlia(*rt, 20, 3, channel)

	go kademlia.network.Listen("localhost", 8000)

	//	kademlia.network.SendPingMessage(&node1)

	//	time.Sleep(5000000000)
	//	kademlia.LookupContact(&node2, 1234)

}

func nextNode(kademliaIDName *KademliaID, kademliaIDNameNext *KademliaID, port int, portNext int) {
	mySelf := NewContact(kademliaIDName, "localhost:"+strconv.Itoa(8000+port))
	next := NewContact(kademliaIDNameNext, "localhost:"+strconv.Itoa(8000+portNext))

	rt := NewRoutingTable(mySelf)

	rt.AddContact(mySelf)
	rt.AddContact(next)

	channel := make(chan []Contact)
	kademlia := NewKademlia(*rt, 20, 3, channel)

	go kademlia.network.Listen("localhost", 8000+port)
}
func finalNode() {
	mySelf := NewContact(NewKademliaID("0000000000000000000000000000000000000099"), "localhost:8099")

	rt := NewRoutingTable(mySelf)

	rt.AddContact(mySelf)

	channel := make(chan []Contact)
	kademlia := NewKademlia(*rt, 20, 3, channel)

	go kademlia.network.Listen("localhost", 8099)

	//	kademlia.network.SendPingMessage(&node1)

	//	time.Sleep(5000000000)
	//	kademlia.LookupContact(&node2, 1234)
}
