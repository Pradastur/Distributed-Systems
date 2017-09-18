package main

type Kademlia struct {
	routingTable RoutingTable
	k            int
	alpha        int
	network      Network
	//data         map[string]File
}

func NewKademlia(rt RoutingTable, k int, alpha int) *Kademlia {
	kademlia := &Kademlia{}
	kademlia.routingTable = rt
	kademlia.k = k
	kademlia.alpha = alpha
	kademlia.network = Network{rt}
	//kademlia.data = make(map[string]File)
	return kademlia
}

func (kademlia *Kademlia) LookupContact(target *Contact) {
	// TODO
}

func (kademlia *Kademlia) LookupData(hash string) {
	// TODO
}

func (kademlia *Kademlia) Store(data []byte) {
	// TODO
}

func (kademlia *Kademlia) ServerThread(port string) { // puede no ser necesario
	// lanzar network
}
