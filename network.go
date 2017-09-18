package main

import "net"
import "bufio"
import "encoding/json"
import "time"
import "fmt"

type Network struct {
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
    Source Contact
    MessageType MessageType
    Content string
}

//UDP

func Listen(ip string, port int) {
<<<<<<< HEAD
	// TODO ServerThread
=======
  fmt.Println("Deploying server thread on port " + string(port))
  port_int, error := strconv.Atoi(port)
  if error != nil {
    // handle error
  }

  // we load the ip for the socket
  addr := net.UDPAddr{
    Port: port,
    IP:   net.ParseIP(ip),
  }

  // create the connection
  conn, err := net.ListenUDP("udp", &addr)
  defer conn.Close()
  if err != nil {
    fmt.Println("No se pudo poner el listen.")
    panic(err)
  }

  for {
    // blocking operation to wait for a message
    fmt.Println("Waiting for inputs... ")
    message, _ := bufio.NewReader(conn).ReadString('\n')


// output message received
    fmt.Println("Server receives: ", string(message))

  }
>>>>>>> 4922648e959a6d3672386f419170e11a4468f7bf
}

func (network *Network) SendPingMessage(contact *Contact) {
	// TODO Client
}

func (network *Network) SendFindContactMessage(contact *Contact) {
	// TODO Client
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO Sprint 2
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO Sprint 2
}
