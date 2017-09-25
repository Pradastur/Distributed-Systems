package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

type Kademlia struct {
	routingTable           RoutingTable
	k                      int
	alpha                  int
	network                Network
	dht                    DHT
	data                   savedata
	alreadyCheckedContacts []Contact
	//data         map[string]File
}

func NewKademlia(rt RoutingTable, k int, alpha int, channel chan []Contact) *Kademlia {
	idMap := make(map[int]bool)
	kademlia := &Kademlia{}
	kademlia.routingTable = rt
	kademlia.k = k
	kademlia.alpha = alpha
	networkMessageRecord := make(map[int]messageControl)
	kademlia.dht = DHT{make(map[KademliaID][]Contact)}
	kademlia.data = savedata{}
	kademlia.network = Network{rt, kademlia, idMap, networkMessageRecord, channel}
	//kademlia.data = make(map[string]File)
	return kademlia
}

func (kademlia *Kademlia) LookupContact(target *Contact, messageID int) {
	//	nilContact := Contact{nil, "", nil}
	//	fmt.Println(kademlia.routingTable.GetContact(*target))
	if kademlia.routingTable.GetContact(*target) != *target {
		contacts := kademlia.routingTable.FindClosestContacts(target.ID, kademlia.alpha)
		kademlia.network.messageRecord[messageID] = messageControl{0, target.ID}
		for i := range contacts {
			fmt.Println("LOOKUP sent to this contact: " + contacts[i].Address)
			kademlia.network.SendFindContactMessage(&contacts[i], target.ID, messageID)
			kademlia.alreadyCheckedContacts = append(kademlia.alreadyCheckedContacts, contacts[i])
		}
	}
	fmt.Println("I already have it")
}

func (kademlia *Kademlia) LookupData(hash string, messageId int) {
	kademliaIdHash := NewKademliaID(hash)
	kademlia.network.messageRecord[messageId] = messageControl{0, kademliaIdHash}
	contact := Contact{kademliaIdHash, "", nil}
	var contactList []Contact
	if kademlia.dht.hasContactsFor(*kademliaIdHash) {
		fmt.Println("LookupData() : Hash is in DHT")
		kademlia.LookupContact(&contact, RandomInt())
		contactList = <-kademlia.network.channel
		for i := 0; i < len(contactList); i++ {
			kademlia.network.SendFindDataMessage(contactList[i], kademliaIdHash, messageId)
		}
	} else {
		fmt.Println("LookupData() : Hash is not in DHT")
		list := kademlia.dht.dhtMap[*kademliaIdHash]
		for i := range list {
			kademlia.network.SendFindDataMessage(list[i], kademliaIdHash, messageId)
		}
	}
}

func (kademlia *Kademlia) Store(data []byte) {
	hash := NewKademliaID(Hash(data))
	kademlia.data.Save(data)
	fmt.Println("Stored hash : " + hash.String())
	pseudoContact := Contact{hash, "", nil}
	kademlia.dht.Update(hash, kademlia.routingTable.me)
	kademlia.LookupContact(&pseudoContact, RandomInt())

	var contactList []Contact // wait for lookup contact response
	contactList = <-kademlia.network.channel
	for i := 0; i < len(contactList); i++ {
		kademlia.network.SendUpdateDHTMessage(contactList[i], hash)
	}
	fmt.Println("Data stored and DHTs updated")
}

func Hash(data []byte) string {
	hashBytes := sha1.Sum(data) // unless kademlia id changes size, this is fine
	return hex.EncodeToString(hashBytes[:])
}
