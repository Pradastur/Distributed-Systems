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

func Listen(ip string, port int) {
	// TODO
}

func SendPingMessage(srcContact Contact, destContact Contact) bool {
  messageToSend := &Message{srcContact, PING, destContact.ID.String()}
  conn, conErr := net.Dial("udp", destContact.Address)
  //jsonToSend, _ := json.Marshal(messageToSend)
  //conn.Write([]byte(jsonToSend))
  fmt.Println(messageToSend)
  //listen for reply
  input := make(chan string, 1)


  if(conErr==nil){
    //fmt.Println("Text to send: ")
    text, err := json.Marshal(messageToSend)

    if (err != nil) {
      fmt.Println("Error ")
      fmt.Println(err)
    }
    //fmt.Println(time.Now().String() + "Message to send server: "+string(text))

    // send to socket
    fmt.Fprintf(conn, string(text) + "\n")
  }
  go getInput(input, conn)

  for {
  		select {
    		case i := <-input:
    			var message Message
    			json.Unmarshal([]byte(i),&message)
    			if(message.MessageType==RESPONSE){
    				return true
    			}
    		case <-time.After(4000 * time.Millisecond):
    			fmt.Println("timed out")
  			  return false
  		}
	}
}

//func getInput(input chan string, destContact Contact, srcContact Contact, conn connection) {
  func getInput(input chan string, conn net.Conn) {
    //for {
    JSONmessage, err := bufio.NewReader(conn).ReadString('\n')

    fmt.Println("Msg rcvd: " + JSONmessage)
    if err != nil {
    }
    input <- JSONmessage
    //}
}

func SendFindContactMessage(srcContact Contact, destContact Contact, contactToFind Contact) []Contact {
	messageToSend := &Message{srcContact, FINDCONTACT,contactToFind.ID.String()}
	//fmt.Print("messageToSend to messageToSend server: "+messageToSend.Content )

	conn, _ := net.Dial("udp", destContact.Address)
	//	  fmt.Print("Text to send: ")
	  text, err := json.Marshal(messageToSend)
	  if err != nil {
		fmt.Println("Error " )
		fmt.Println(err)
	}
	  //fmt.Println("Message to send server: "+string(text))

	  // send to socket
	  fmt.Fprintf(conn, string(text) + "\n")
	  // listen for reply
	  JSONmessage, _ := bufio.NewReader(conn).ReadString('\n')
	  var message Message
	  json.Unmarshal([]byte(JSONmessage),&message)
	  var contacts []Contact
	  json.Unmarshal([]byte(message.Content),&contacts)
	  /*for i := range contacts {
	  	fmt.Println("Message from server " +string(i) +" : "+ contacts[i].ID.String())
	  }*/
	  return contacts
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO Sprint2
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO Sprint2
}
