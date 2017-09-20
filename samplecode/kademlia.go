package main

import "fmt"

type Kademlia struct {
	routingTable RoutingTable
	k            int
	alpha        int
	network      Network
	//data         map[string]File
}

func NewKademlia(rt RoutingTable, k int, alpha int, channel chan []Contact) *Kademlia {
	var IdArray []int
	kademlia := &Kademlia{}
	kademlia.routingTable = rt
	kademlia.k = k
	kademlia.alpha = alpha
	networkMessageMap := make(map[int]Contact)
	networkMessageRecord := make(map[int]messageControl)
	kademlia.network = Network{rt, kademlia, IdArray, networkMessageRecord, networkMessageMap, channel}
	//kademlia.data = make(map[string]File)
	return kademlia
}

func (kademlia *Kademlia) LookupContact(target *Contact, messageID int) {
	contacts := kademlia.routingTable.FindClosestContacts(target.ID, kademlia.alpha)
	kademlia.network.messageRecord[messageID] = messageControl{0, target.ID}
	//fmt.Println("SOURCE: " + source.Address)
	for i := range contacts {
		fmt.Println("LOOKUP sent to this contact: " + contacts[i].Address)
		//if source.Address != contacts[i].Address {
		kademlia.network.SendFindContactMessage(&contacts[i], target.ID, messageID)
		//}
	}
}

func (kademlia *Kademlia) LookupData(hash string) {

}

func (kademlia *Kademlia) Store(data []byte) {
	// TODO Sprint 2
}
