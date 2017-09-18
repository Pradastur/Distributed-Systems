package d7024e

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
	// TODO ServerThread
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
