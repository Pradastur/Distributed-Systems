package main

import (
	"fmt"
	"testing"
)

func TestRoutingTable(t *testing.T) {
	srcContact := NewRoutingTable(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))
	rt := NewRoutingTable(NewContact(NewKademliaID(srcContact)

	rt.AddContact(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
	rt.AddContact(NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8003"))
	rt.AddContact(NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8004"))
	rt.AddContact(NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8005"))
	rt.AddContact(NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8006"))

	contacts := rt.FindClosestContacts(NewKademliaID("2111111400000000000000000000000000000000"), 20)
	for i := range contacts {
		fmt.Println(contacts[i].String())
	}

	kademlia:=NewKademlia(*rt,2,1)
	go kademlia.ServerThread("8000")
<<<<<<< HEAD
=======
	fmt.Println(SendPingMessage(srcContact, srcContact))
>>>>>>> b799241e9953bb156fdedc5e994fabaecb300906
}
