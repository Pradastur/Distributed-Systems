package d7024e

import (
	"fmt"
	"testing"
)

func TestRoutingTable(t *testing.T) {
	rt := NewRoutingTable(NewContact(NewKademliaID("8000000000000000000000000000000000000000"), "localhost:8000"))

	rt.AddContact(NewContact(NewKademliaID("7000000000000000000000000000000000000000"), "localhost:8001"))
	rt.AddContact(NewContact(NewKademliaID("6000000000000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("5000000000000000000000000000000000000000"), "localhost:8003"))
	rt.AddContact(NewContact(NewKademliaID("4000000000000000000000000000000000000000"), "localhost:8004"))
	rt.AddContact(NewContact(NewKademliaID("3000000000000000000000000000000000000000"), "localhost:8005"))
	rt.AddContact(NewContact(NewKademliaID("2000000000000000000000000000000000000000"), "localhost:8006"))
	rt.AddContact(NewContact(NewKademliaID("9000000000000000000000000000000000000000"), "localhost:8007"))
	rt.AddContact(NewContact(NewKademliaID("F000000000000000000000000000000000000000"), "localhost:8008"))

	contacts := rt.FindClosestContacts(NewKademliaID("1000000000000000000000000000000000000000"), 20)
	for i := range contacts {
		fmt.Println(contacts[i].String())
	}
}
