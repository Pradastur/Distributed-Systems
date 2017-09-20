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
	messageIDs    []int
	messageRecord map[int]messageControl
	messageMap    map[int]Contact
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
	R_STORE
)

type Message struct {
	Source      Contact
	MessageType MessageType
	Content     string
	ID          int
}

func (network *Network) checkMessage(messageID int) bool {
	proceed := true
	for i := 0; i < len(network.messageIDs); i++ {
		if network.messageIDs[i] == messageID {
			proceed = false
		}
	}

	if (network.messageMap[messageID] != Contact{nil, "", nil}) {
		proceed = true
		//network.messageMap[messageID] = Contact{nil, "", nil}
	}

	return proceed
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
		fmt.Println("Waiting for inputs... (SERVER " + port_string + ")")
		message, _ := bufio.NewReader(conn).ReadString('\n')

		// output message received
		//fmt.Println("(SERVER "+port_string+") receives: ", string(message))

		// EVALUAR LO K LLEGA
		var messageDecoded Message
		json.Unmarshal([]byte(message), &messageDecoded)

		//fmt.Println("(SERVER " + port_string + ") receives message type: ", string(messageDecoded.Content[1]))
		network.routingTable.AddContact(messageDecoded.Source) //source is added to the contacts
		fmt.Println("(SERVER "+port_string+") receives message type: ", messageDecoded.MessageType)

		if network.checkMessage(messageDecoded.ID) {
			switch messageDecoded.MessageType {
			case PING:
				//fmt.Println("Message Ping Received:", string(messageDecoded.Content[0]))
				// go kademlia.routingTable.AddContact(messageDecoded.Source)
				//network.routingTable.AddContact(messageDecoded.Source)
				network.SendMessage(&messageDecoded.Source, RESPONSE, "", -1)
				fmt.Println("PING")
				break

			case FIND_NODE:
				//fmt.Println("We're looking for the contact (FIND_NODE)")
				wanted := Contact{NewKademliaID(messageDecoded.Content), "", nil}
				fmt.Println("THE WANTED NODE: " + wanted.String())
				contacts := network.kademlia.routingTable.FindClosestContacts(wanted.ID, network.kademlia.k)
				contactsMarshalized, _ := json.Marshal(contacts)
				network.SendMessage(&messageDecoded.Source, R_FIND_NODE, string(contactsMarshalized), messageDecoded.ID)

				/*contactInRT := network.routingTable.GetContact(wanted) // contact
				if contactInRT.ID != nil {
					wantedMarshalized, _ := json.Marshal(contactInRT)
					network.SendMessage(&messageDecoded.Source, R_FIND_NODE, string(wantedMarshalized), messageDecoded.ID)
				} else {
					fmt.Println("Entering LOOKUP, (node " + port_string + ") save address " + messageDecoded.Source.Address)
					network.messageMap[messageDecoded.ID] = messageDecoded.Source
					fmt.Println(network.messageMap[messageDecoded.ID])
					network.kademlia.LookupContact(&wanted, messageDecoded.ID)
				}*/

				break

			case R_FIND_NODE:
				if network.messageRecord[messageDecoded.ID].wanted != nil {
					fmt.Println("CURRENT MESSAGE RECORD :")
					fmt.Println(network.messageRecord[messageDecoded.ID])
					network.messageRecord[messageDecoded.ID] = messageControl{network.messageRecord[messageDecoded.ID].count + 1, network.messageRecord[messageDecoded.ID].wanted}
					control := network.messageRecord[messageDecoded.ID]
					fmt.Println("COUNT VALUE :")
					fmt.Println(control.count)
					if control.count < network.kademlia.k-1 {
						var contactList []Contact
						json.Unmarshal([]byte(messageDecoded.Content), &contactList)
						contactCount := len(contactList)

						var maxIndex int
						if contactCount < network.kademlia.alpha {
							maxIndex = contactCount
						} else {
							maxIndex = network.kademlia.alpha
						}

						localContacts := network.routingTable.FindClosestContacts(control.wanted, network.kademlia.alpha)
						nilContact := Contact{nil, "", nil}
						for i := 0; i < len(contactList); i++ {
							for j := 0; j < len(localContacts); j++ {
								if contactList[i] == localContacts[j] {
									contactList[i] = nilContact
								}
							}
						}

						fmt.Println("CLOSEST CONTACTS:")
						fmt.Println(localContacts)

						for i := 0; i < maxIndex; i++ {
							if contactList[i] != nilContact {
								network.routingTable.AddContact(contactList[i])
								network.SendFindContactMessage(&contactList[i], control.wanted, messageDecoded.ID)
							}
						}
					} else {
						network.channel <- network.routingTable.FindClosestContacts(control.wanted, network.kademlia.alpha)
					}
				}
				//network.routingTable.AddContact(contactFound)

				// if we have to send smth back
				/*var contactInMap Contact
				contactInMap = network.messageMap[messageDecoded.ID]
				fmt.Println("Print CONTACT IN MAP of R_FIND_NODE (node " + port_string + " )")
				fmt.Println(contactInMap)
				if (contactInMap != Contact{nil, "", nil}) {
					fmt.Println("HOLA ESTAMOS LLEGANDO AQUI?")
					network.SendMessage(&contactInMap, R_FIND_NODE, messageDecoded.Content, messageDecoded.ID)
					//network.messageMap[messageDecoded.ID] = Contact{nil, "", nil}
				}*/
				break

			case FIND_VALUE:
				break

			case RESPONSE:
				fmt.Println("RESPONSE")
				break
			}
		} else {
			fmt.Println("IGNORANDO MENSAJE - ID REPETIDA")
		}
	}
}

func (network *Network) SendMessage(contact *Contact, mType MessageType, content string, messageID int) {
	ID := rand.Intn(999999999)
	if messageID != -1 {
		ID = messageID
	}
	messageToSend := &Message{network.routingTable.me, mType, content, ID}
	conn, conErr := net.Dial("udp", contact.Address)
	if conErr != nil {
		fmt.Println("No se puede crear la conexion de salida (CLIENTE).")
	}

	text, _ := json.Marshal(messageToSend)
	fmt.Println("Message: " + string(text) + " Dest: " + contact.Address)

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

func (network *Network) SendFindDataMessage(hash string) {
	// SPRINT 2
}

func (network *Network) SendStoreMessage(data []byte) {
	// SPRINT 2
}
