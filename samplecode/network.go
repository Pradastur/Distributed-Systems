package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"strconv"
)

type Network struct {
	routingTable RoutingTable
	kademlia     *Kademlia
	messageIDs   []int
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

func (network *Network) checkearMessage(messageID int) bool {
	for i := 0; i < len(network.messageIDs); i++ {
		if network.messageIDs[i] == messageID {
			return true
		}
		if network.messageIDs[i] != messageID {
			return false
		}
	}
	return false
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
		fmt.Println("(SERVER "+port_string+") receives message type: ", messageDecoded.MessageType)
		network.routingTable.AddContact(messageDecoded.Source) //source is added to the contacts
		switch messageDecoded.MessageType {
		case PING:
			//fmt.Println("Message Ping Received:", string(messageDecoded.Content[0]))
			// go kademlia.routingTable.AddContact(messageDecoded.Source)
			//network.routingTable.AddContact(messageDecoded.Source)
			network.SendMessage(&messageDecoded.Source, RESPONSE, "")
			fmt.Println("PING")
			break

		case FIND_NODE:
			//fmt.Println("We're looking for the contact (FIND_NODE)")
			wanted := Contact{NewKademliaID(messageDecoded.Content), "", nil}

			contactInRT := network.routingTable.GetContact(wanted) // contact
			if contactInRT.ID != nil {
				wantedMarshalized, _ := json.Marshal(contactInRT)
				network.SendMessage(&messageDecoded.Source, R_FIND_NODE, string(wantedMarshalized))
			} else {
				network.kademlia.LookupContact(&wanted)
			}

			break

		case R_FIND_NODE:
			var contactFound Contact
			json.Unmarshal([]byte(messageDecoded.Content), &contactFound)

			network.routingTable.AddContact(contactFound)
			break

		case FIND_VALUE:
			break

		case RESPONSE:
			fmt.Println("RESPONSE")
			break
		}

	}
}

func (network *Network) SendMessage(contact *Contact, mType MessageType, content string) {
	ID := rand.Intn(999999999)
	messageToSend := &Message{network.routingTable.me, mType, content, ID}
	conn, conErr := net.Dial("udp", contact.Address)
	if conErr != nil {
		fmt.Println("No se puede crear la conexion de salida (CLIENTE).")
	}

	text, _ := json.Marshal(messageToSend)
	fmt.Println("Message to send server: " + string(text))

	// send to socket
	fmt.Fprintf(conn, string(text)+"\n")
	network.messageIDs = append(network.messageIDs, ID)
}

func (network *Network) SendPingMessage(contact *Contact) {
	network.SendMessage(contact, PING, "")
}

func (network *Network) SendFindContactMessage(contact *Contact, targetID *KademliaID) {
	network.SendMessage(contact, FIND_NODE, targetID.String())
}

func (network *Network) SendFindDataMessage(hash string) {
	// SPRINT 2
}

func (network *Network) SendStoreMessage(data []byte) {
	// SPRINT 2
}
