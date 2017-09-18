package main

import (
	"fmt"
	"testing"
)

func TestRoutingTable(t *testing.T) {
	go nodoUno()
	go nodoDos()
	for {
	}
	//fmt.Println(SendPingMessage(srcContact, srcContact))
}

func nodoUno() {
	srcContact := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000")
	rt := NewRoutingTable(srcContact)

	destino := NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002")

	rt.AddContact(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
	rt.AddContact(destino)
	rt.AddContact(NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002"))

	contacts := rt.FindClosestContacts(NewKademliaID("2111111400000000000000000000000000000000"), 20)
	for i := range contacts {
		fmt.Println(contacts[i].String())
	}

	kademlia := NewKademlia(*rt, 2, 1)
	go kademlia.network.Listen("localhost", 8000)
	//kademlia.network.SendPingMessage(&destino)
}

func nodoDos() {
	srcContact := NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002")
	rt := NewRoutingTable(srcContact)

	destino := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000")

	rt.AddContact(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
	rt.AddContact(destino)
	rt.AddContact(NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002"))

	contacts := rt.FindClosestContacts(NewKademliaID("2111111400000000000000000000000000000000"), 20)
	for i := range contacts {
		fmt.Println(contacts[i].String())
	}

	kademlia := NewKademlia(*rt, 2, 1)
	go kademlia.network.Listen("localhost", 8002)
	kademlia.network.SendPingMessage(&destino)
}
