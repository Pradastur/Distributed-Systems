<<<<<<< HEAD
package main 
=======
package main
>>>>>>> 767f8db8bd364d4c57f29672aff47286b89d27ee

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

<<<<<<< HEAD

// used http://technosophos.com/2014/03/18/a-simple-udp-server-in-go.html
func (kademlia *Kademlia) ServerThread(port string) {
  fmt.Println("Deploying server thread on port " + port)
  port_int, error := strconv.Atoi(por











      if error != nil {
    // handle error
  }

  // we load the ip for the socket
  addr := net.UDPAddr{
    Port: port_int,
    IP:   net.ParseIP("localhost:8000"),
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

    var messageDecoded Message
    json.Unmarshal([]byte(message), &messageDecoded)
    //fmt.Println("Message type Received:", messageDecoded.MessageType)

    var responseMessage Message

    //fmt.Println("Message Ping Received:", string(messageDecoded.Content[0]))
    //go kademlia.routingTable.AddContact(messageDecoded.Source)
    responseMessage = Message{kademlia.routingTable.me, RESPONSE, ""}
    JSONResponseMessage, _ := json.Marshal(responseMessage)
    // sample process for string received
    //var a []byte = []byte("Response \n");

    //fmt.Print("message to byte", string(JSONResponseMessage))
    //conn.Write([]byte(string(JSONResponseMessage) + "\n"))
		fmt.Fprintf(conn, string(JSONResponseMessage) + "\n")
		fmt.Println("Json Response: " + string(JSONResponseMessage))
  }
=======
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
<<<<<<< HEAD
>>>>>>> 767f8db8bd364d4c57f29672aff47286b89d27ee
=======
	fmt.Println("I already have it")
>>>>>>> bc56a38fbdd1e2442556219978fe5ecc127d9b47
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
