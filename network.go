package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"strconv"
)

type messageControl struct {
	count  int
	wanted *KademliaID
}

type Network struct {
	routingTable  RoutingTable
	kademlia      *Kademlia
	messageIDs    map[int]bool
	messageRecord map[int]messageControl
	channel       chan []Contact
}

type MessageType int

const (
	PING MessageType = 1 + iota
	RESPONSE

	FIND_NODE
	R_FIND_NODE
	FIND_VALUE
	R_FIND_VALUE
	STORE
	UPDATE_DHT
)

type Message struct {
	Source      Contact
	MessageType MessageType
	Content     string
	ID          int
}

func (network *Network) checkMessage(messageID int) bool {
	fmt.Println("Received MessageId :" + strconv.Itoa(messageID))
	for msgID, proceed := range network.messageIDs {
		if msgID == messageID {
			return proceed
		}
	}
	network.messageIDs[messageID] = false
	return false
}

func (network *Network) setProceed(messageID int) {
	network.messageIDs[messageID] = true
}

func (network *Network) Listen(ip string, port int) {
	port_string := strconv.Itoa(port)

	fmt.Println("Deploying SERVER thread on port " + port_string)

	// we load the ip for the socket
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP(ip),
	}

	// create the connection
	conn, err := net.ListenUDP("udp", &addr)
	defer conn.Close()
	if err != nil {
		fmt.Println("No se pudo poner el listen. (SERVER " + port_string + ")")
		panic(err)
	}
	for {
		// blocking operation to wait for a message
		fmt.Println("(SERVER " + port_string + ") Waiting for inputs...")
		message, _ := bufio.NewReader(conn).ReadString('\n')

		// output message received
		//fmt.Println("(SERVER "+port_string+") receives: ", string(message))

		// EVALUAR LO K LLEGA
		var messageDecoded Message
		json.Unmarshal([]byte(message), &messageDecoded)

		//fmt.Println("(SERVER " + port_string + ") receives message type: ", string(messageDecoded.Content[1]))
		network.routingTable.AddContact(messageDecoded.Source) //source is added to the contacts
		fmt.Println("(SERVER "+port_string+") receives message type: ", messageDecoded.MessageType)

		if !network.checkMessage(messageDecoded.ID) {
			switch messageDecoded.MessageType {
			case PING:
				//fmt.Println("Message Ping Received:", string(messageDecoded.Content[0]))
				// go kademlia.routingTable.AddContact(messageDecoded.Source)
				//network.routingTable.AddContact(messageDecoded.Source)
				network.SendMessage(&messageDecoded.Source, RESPONSE, "", -1)
				fmt.Println("PING")
				network.setProceed(messageDecoded.ID)
				break

			case FIND_NODE:
				//fmt.Println("We're looking for the contact (FIND_NODE)")
				wanted := Contact{NewKademliaID(messageDecoded.Content), "", nil}
				fmt.Println("THE WANTED NODE: " + wanted.String())
				contacts := network.kademlia.routingTable.FindClosestContacts(wanted.ID, network.kademlia.k)
				contactsMarshalized, _ := json.Marshal(contacts)
				network.SendMessage(&messageDecoded.Source, R_FIND_NODE, string(contactsMarshalized), messageDecoded.ID)
				network.setProceed(messageDecoded.ID)
				break

			case R_FIND_NODE:
				fmt.Println("(SERVER " + port_string + ") R_FIND_NODE received")

				if network.messageRecord[messageDecoded.ID].wanted != nil {
					// fmt.Println("(SERVER " + port_string + ") R_FIND_NODE : inside first if")
					network.messageRecord[messageDecoded.ID] = messageControl{network.messageRecord[messageDecoded.ID].count + 1, network.messageRecord[messageDecoded.ID].wanted}
					control := network.messageRecord[messageDecoded.ID]
					if control.count < network.kademlia.k-1 {
						// fmt.Println("(SERVER " + port_string + ") R_FIND_NODE : count inferior to k-1")
						var contactList []Contact
						json.Unmarshal([]byte(messageDecoded.Content), &contactList)
						contactCount := len(contactList)

						var maxIndex int
						if contactCount < network.kademlia.alpha {
							maxIndex = contactCount
						} else {
							maxIndex = network.kademlia.alpha
						}

						var alreadyChecked bool
						var hasSendSomething bool
						hasSendSomething = false
						for i := 0; i < maxIndex; i++ {
							alreadyChecked = false
							for j := range network.kademlia.alreadyCheckedContacts {
								if contactList[i].ID.Equals(network.kademlia.alreadyCheckedContacts[j].ID) {
									alreadyChecked = true
								}
							}
							if !alreadyChecked {
								network.routingTable.AddContact(contactList[i])
								network.SendFindContactMessage(&contactList[i], control.wanted, messageDecoded.ID)
								network.kademlia.alreadyCheckedContacts = append(network.kademlia.alreadyCheckedContacts, contactList[i])
								fmt.Println("(SERVER " + port_string + ") sent messageID " + strconv.Itoa(messageDecoded.ID) + " to " + contactList[i].Address)
								hasSendSomething = true
							}
						}
						if !hasSendSomething {
							contactList := network.routingTable.FindClosestContacts(control.wanted, network.kademlia.alpha)
							network.channel <- contactList

							fmt.Println("R_FIND_NODE : LookupContact is done, transmitted results : ")
							fmt.Println(contactList)

							network.setProceed(messageDecoded.ID)
						}
					} else {
						if control.count == network.kademlia.k-1 {
							contactList := network.routingTable.FindClosestContacts(control.wanted, network.kademlia.alpha)
							network.channel <- contactList

							fmt.Println("R_FIND_NODE : LookupContact is done, transmitted results : ")
							fmt.Println(contactList)

							network.setProceed(messageDecoded.ID)
						}
					}
				}

				break

			case FIND_VALUE:
				hash := messageDecoded.Content
				ourFileHash := Hash(network.kademlia.data.Get())

				fmt.Println(hash)
				fmt.Println(ourFileHash)

				if hash == ourFileHash { // I HAVE IT
					fmt.Println("FIND_VALUE : I have the data case")
					network.SendStoreMessage(messageDecoded.Source, network.kademlia.data.Get())
				} else { // I DONT HAVE IT
					if network.kademlia.dht.hasContactsFor(*NewKademliaID(hash)) {
						fmt.Println("FIND_VALUE : I have the hash in DHT")
						contactList := network.kademlia.dht.dhtMap[*NewKademliaID(hash)]
						contactListMarshalized, _ := json.Marshal(contactList)
						network.SendMessage(&messageDecoded.Source, R_FIND_VALUE, string(contactListMarshalized), messageDecoded.ID)
					} else {
						fmt.Println("FIND_VALUE : I have nothing")
						contactList := network.routingTable.FindClosestContacts(NewKademliaID(hash), network.kademlia.alpha)
						contactListMarshalized, _ := json.Marshal(contactList)
						network.SendMessage(&messageDecoded.Source, R_FIND_VALUE, string(contactListMarshalized), messageDecoded.ID)
					}
				}
				network.setProceed(messageDecoded.ID)
				break

			case R_FIND_VALUE:
				fmt.Println("R_FIND_VALUE : just received")
				if network.messageRecord[messageDecoded.ID].wanted != nil {
					network.messageRecord[messageDecoded.ID] = messageControl{network.messageRecord[messageDecoded.ID].count + 1, network.messageRecord[messageDecoded.ID].wanted}
					control := network.messageRecord[messageDecoded.ID]
					if control.count < network.kademlia.k-1 {
						var contactList []Contact
						json.Unmarshal([]byte(messageDecoded.Content), &contactList)
						for i := 0; i < len(contactList); i++ {
							network.SendFindDataMessage(contactList[i], network.messageRecord[messageDecoded.ID].wanted, messageDecoded.ID)
						}
					} else {
						fmt.Println("R_FIND_VALUE : count is bigger than k. Data doesn't exist")
						network.setProceed(messageDecoded.ID)
					}
				}

				break

			case STORE:
				var newData []byte
				json.Unmarshal([]byte(newData), messageDecoded.Content)
				network.kademlia.data.Save(newData)
				fmt.Println("STORE MSG RECEIVED, File SAVED")
				hash := NewKademliaID(Hash(newData))
				network.kademlia.dht.Update(hash, network.kademlia.routingTable.me)

				ct := Contact{hash, "", nil}
				network.kademlia.LookupContact(&ct, RandomInt())
				go network.StoreLookupContactCallback(hash)
				network.setProceed(messageDecoded.ID)
				break

			case UPDATE_DHT:
				hash := NewKademliaID(messageDecoded.Content)
				network.kademlia.dht.Update(hash, messageDecoded.Source)
				network.setProceed(messageDecoded.ID)
				break
			}
		} else {
			fmt.Println("IGNORANDO MENSAJE - ID REPETIDA")
		}
	}
}

func (network *Network) SendMessage(contact *Contact, mType MessageType, content string, messageID int) {
	ID := RandomInt()
	if messageID != -1 {
		ID = messageID
	}
	messageToSend := &Message{network.routingTable.me, mType, content, ID}
	conn, conErr := net.Dial("udp", contact.Address)
	if conErr != nil {
		fmt.Println("No se puede crear la conexion de salida (CLIENTE).")
	}

	text, _ := json.Marshal(messageToSend)
	// fmt.Println("Message: " + string(text) + " Dest: " + contact.Address)

	// send to socket
	fmt.Fprintf(conn, string(text)+"\n")
}

func (network *Network) SendPingMessage(contact *Contact) {
	network.SendMessage(contact, PING, "", -1)
}

func (network *Network) SendFindContactMessage(contact *Contact, targetID *KademliaID, messageID int) {
	if contact.Address != network.kademlia.routingTable.me.Address {
		network.SendMessage(contact, FIND_NODE, targetID.String(), messageID)
	}
}

func (network *Network) SendFindDataMessage(contact Contact, hash *KademliaID, messageID int) {
	if messageID == 0 {
		messageID = RandomInt()
	}

	network.SendMessage(&contact, FIND_VALUE, hash.String(), messageID)
	fmt.Println("SendFindDataMessage() done")
}

func (network *Network) SendStoreMessage(contact Contact, data []byte) {
	content, _ := json.Marshal(data)
	network.SendMessage(&contact, STORE, string(content), RandomInt())
}

func (network *Network) SendUpdateDHTMessage(contact Contact, hash *KademliaID) {
	network.SendMessage(&contact, UPDATE_DHT, hash.String(), RandomInt())
}

func (network *Network) StoreLookupContactCallback(hash *KademliaID) {
	fmt.Println("STORE MSG : Now waiting for LookupContact results")
	var contactList []Contact // wait for lookup contact response
	contactList = <-network.channel
	fmt.Println("LOOKUP CONTACT CALLBACK RECEIVED")
	for i := range contactList {
		network.SendUpdateDHTMessage(contactList[i], hash)
	}
}

func RandomInt() int {
	return rand.Intn(999999999)
}
