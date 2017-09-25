package main

type DHT struct {
	dhtMap map[KademliaID][]Contact
}

func (dht *DHT) Update(hash *KademliaID, contact Contact) {
	contactList := dht.dhtMap[*hash]
	dht.dhtMap[*hash] = append(contactList, contact)
}

func (dht *DHT) hasContactsFor(hash KademliaID) bool {
	contactList := dht.dhtMap[hash]
	return len(contactList) > 0
}
