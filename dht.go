package main

type DHT struct {
	dhtMap map[KademliaID][]Contact
}

func (dht *DHT) Update(hash *KademliaID, contact Contact) {
	contactList := dht.dhtMap[*hash]
	dht.dhtMap[*hash] = append(contactList, contact)

	/*if contactList != nil {
	    append(contactList, contact)
	    dht.dhtMap[hash] = contactList
		}else{
	    dht.dhtMap[hash] =
	  }*/
}

func (dht *DHT) hasContactsFor(hash KademliaID) bool {
	contactList := dht.dhtMap[hash]
	return len(contactList) > 0
}
