package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"time"
)

type Network struct {
	routingTable RoutingTable
}

type MessageType int

const (
	PING MessageType = 1 + iota
	FINDCONTACT
	FINDDATA
	STORE
	ADDNODE
	RESPONSE
)

type Message struct {
	Source      Contact
	MessageType MessageType
	Content     string
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
		fmt.Println("(SERVER "+port_string+") receives: ", string(message))

		// EVALUAR LO K LLEGA
	}
}

func (network *Network) SendMessage(contact *Contact, mType MessageType, content string) {
	messageToSend := &Message{network.routingTable.me, mType, content}
	conn, conErr := net.Dial("udp", contact.Address)
	if conErr != nil {
		fmt.Println("No se puede crear la conexion de salida (CLIENTE).")
	}

	text, _ := json.Marshal(messageToSend)
	fmt.Println(time.Now().String() + "Message to send server: " + string(text))

	// send to socket
	fmt.Fprintf(conn, string(text)+"\n")
}

func (network *Network) SendPingMessage(contact *Contact) {
	network.SendMessage(contact, PING, contact.ID.String())
}

func (network *Network) SendFindContactMessage(contact *Contact) {
	network.SendMessage(contact,FINDCONTACT, contact.ID.String())
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}
