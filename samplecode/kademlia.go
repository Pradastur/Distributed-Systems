package main

type Kademlia struct {
	routingTable RoutingTable
	k            int
	alpha        int
	network      Network
	//data         map[string]File
}

func NewKademlia(rt RoutingTable, k int, alpha int) *Kademlia {
	var IdArray []int
	kademlia := &Kademlia{}
	kademlia.routingTable = rt
	kademlia.k = k
	kademlia.alpha = alpha
	kademlia.network = Network{rt, kademlia, IdArray}
	//kademlia.data = make(map[string]File)
	return kademlia
}

func (kademlia *Kademlia) LookupContact(target *Contact) {
	contacts := kademlia.routingTable.FindClosestContacts(target.ID, kademlia.k)
	for i := range contacts {
		kademlia.network.SendFindContactMessage(&contacts[i], target.ID)
	}
}

func (kademlia *Kademlia) LookupData(hash string) {
	// TODO Sprint 2
}

func (kademlia *Kademlia) Store(data []byte) {
	// TODO Sprint 2
}
