package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"sync"
	"time"
)

type Kademlia struct {
	routingTable           RoutingTable
	k                      int
	alpha                  int
	network                Network
	dht                    DHT
	fSys                   Filesystem
	alreadyCheckedContacts []Contact
}

func NewKademlia(rt RoutingTable, k int, alpha int, channel chan []Contact) *Kademlia {
	idMap := make(map[int]bool)
	kademlia := &Kademlia{}
	kademlia.routingTable = rt
	kademlia.k = k
	kademlia.alpha = alpha
	kademlia.dht = DHT{make(map[KademliaID][]Contact)}
	kademlia.fSys = Filesystem{}
	record := MessageRecordMutex{make(map[int]messageControl), sync.Mutex{}}
	rout := RoutingTableMutex{rt, sync.Mutex{}}
	kademlia.network = Network{rout, kademlia, idMap, channel, record}
	return kademlia
}

func (kademlia *Kademlia) LookupContact(target *Contact, messageID int) {
	if kademlia.routingTable.GetContact(*target) != *target {
		contacts := kademlia.routingTable.FindClosestContacts(target.ID, kademlia.alpha)
		for j := range contacts {
			fmt.Println("ESTA ES LA LISTA DE CONTACTOS: " + contacts[j].ID.String())
		}
		kademlia.network.record.mutex.Lock()
		kademlia.network.record.messageRecord[messageID] = messageControl{0, target.ID}
		kademlia.network.record.mutex.Unlock()
		for i := range contacts {
			fmt.Println("LOOKUP sent to this contact: " + contacts[i].Address)
			kademlia.network.SendFindContactMessage(&contacts[i], target.ID, messageID)
			kademlia.alreadyCheckedContacts = append(kademlia.alreadyCheckedContacts, contacts[i])
		}
	}
	fmt.Println("I already have it in my routing table")
}

func (kademlia *Kademlia) LookupData(hash string, messageId int) {

	kademliaIdHash := NewKademliaID(hash)
	kademlia.network.record.mutex.Lock()
	kademlia.network.record.messageRecord[messageId] = messageControl{0, kademliaIdHash}
	kademlia.network.record.mutex.Unlock()
	contact := Contact{kademliaIdHash, "", nil}
	var contactList []Contact
	if kademlia.fSys.hasData(hash) {
		fmt.Println("***********I HAVE IT************")
		fmt.Println(kademlia.fSys.getFile(hash))
		os.Exit(0)
	} else {
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
}

func (kademlia *Kademlia) Store(path string, pinned bool, date time.Time, data []byte) {
	hash := NewKademliaID(Hash(data))
	file := NewFile(path, pinned, data)
	kademlia.fSys.save(file)
	fmt.Println("Stored hash : " + hash.String())
	pseudoContact := Contact{hash, "", nil}
	kademlia.dht.Update(hash, kademlia.routingTable.me)
	fmt.Println("Update")
	kademlia.LookupContact(&pseudoContact, RandomInt())
	fmt.Println("LookkUpContact")

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
