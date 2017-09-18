package main 

import "net"
import "fmt"
import "strconv"
import "bufio"
import "encoding/json"

type Kademlia struct {
  routingTable RoutingTable
  k            int
  alpha        int
  network      Network
}

func NewKademlia(rt RoutingTable, k int, alpha int) *Kademlia {
	kademlia := &Kademlia{}
	kademlia.routingTable = rt
	kademlia.k = k
	kademlia.alpha = alpha
	kademlia.network = Network{}
	return kademlia
}


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
}

func (kademlia *Kademlia) LookupContact(target *Contact) {
	// TODO
}

func (kademlia *Kademlia) LookupData(hash string) {
	// TODO
}

func (kademlia *Kademlia) Store(data []byte) {
	// TODO
}
