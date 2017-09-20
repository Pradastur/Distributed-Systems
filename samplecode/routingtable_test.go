package main

import (
	"fmt"
	"testing"
)

func TestRoutingTable(t *testing.T) {
	go node1()
	go node2()
	go node3()
	go node4()
	for {
	}
	//fmt.Println(SendPingMessage(srcContact, srcContact))
}

func node1() {
	srcContact := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000")
	rt := NewRoutingTable(srcContact)

	destino := NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002")
	unknownContact := NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "")

	rt.AddContact(srcContact) //NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
	rt.AddContact(destino)
	//rt.AddContact(NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002"))
	//rt.AddContact(NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8002"))
	//rt.AddContact(NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8002"))
	//rt.AddContact(NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002"))

	contacts := rt.FindClosestContacts(NewKademliaID("2111111400000000000000000000000000000000"), 20)
	for i := range contacts {
		fmt.Println("N1 Contact: " + contacts[i].String())
	}

	channel := make(chan []Contact)
	kademlia := NewKademlia(*rt, 20, 3, channel)
	fmt.Println(kademlia.network.messageMap[1])

	go kademlia.network.Listen("localhost", 8000)
	//kademlia.network.SendPingMessage(&destino)
	kademlia.LookupContact(&unknownContact, 1234)
	var contactList []Contact

	contactList = <-channel
	fmt.Println("LIST OF CONTACTS RECEIVED WOHOOO (i think lol)")
	fmt.Println(contactList)
}

func node2() {
	srcContact := NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002")
	node3Contact := NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8003")

	rt := NewRoutingTable(srcContact)

	rt.AddContact(srcContact)
	rt.AddContact(node3Contact)
	channel := make(chan []Contact)
	kademlia := NewKademlia(*rt, 20, 3, channel)
	go kademlia.network.Listen("localhost", 8002)
}

func node3() {
	srcContact := NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8003")
	node4Contact := NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8004")
	rt := NewRoutingTable(srcContact)
	rt.AddContact(srcContact)
	rt.AddContact(node4Contact)
	channel := make(chan []Contact)
	kademlia := NewKademlia(*rt, 20, 3, channel)
	go kademlia.network.Listen("localhost", 8003)

}

func node4() {
	srcContact := NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8004")
	rt := NewRoutingTable(srcContact)
	rt.AddContact(srcContact)
	channel := make(chan []Contact)
	kademlia := NewKademlia(*rt, 20, 3, channel)
	go kademlia.network.Listen("localhost", 8004)
}
