package main

import (
	"container/list"
)

type bucket struct {
	list *list.List
}

func newBucket() *bucket {
	bucket := &bucket{}
	bucket.list = list.New()
	return bucket
}

func (bucket *bucket) GetContact(contact Contact) Contact {
	for e := bucket.list.Front(); e != nil; e = e.Next() {
		node := e.Value.(Contact)
		nodeID := node.ID

		if (contact).ID.Equals(nodeID) {
			return node
		}
	}

	return Contact{nil, "", nil}
}

func (bucket *bucket) AddContact(contact Contact) {
	var element *list.Element
	for e := bucket.list.Front(); e != nil; e = e.Next() {
		nodeID := e.Value.(Contact).ID

		if (contact).ID.Equals(nodeID) {
			element = e
		}
	}

	if element == nil {
		if bucket.list.Len() < bucketSize {
			bucket.list.PushFront(contact)
		}
	} else {
		bucket.list.MoveToFront(element)
	}
}

func (bucket *bucket) GetContactAndCalcDistance(target *KademliaID) []Contact {
	var contacts []Contact

	for elt := bucket.list.Front(); elt != nil; elt = elt.Next() {
		contact := elt.Value.(Contact)
		contact.CalcDistance(target)
		contacts = append(contacts, contact)
	}

	return contacts
}

func (bucket *bucket) Len() int {
	return bucket.list.Len()
}

func (bucket *bucket) String() string{
	output := ""
	for e := bucket.list.Front(); e != nil; e = e.Next() {
	 output += e.Value.(Contact).Address
	}
	return output
}
