package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("")
	fmt.Println("WELCOME TO KADEMLIA")
	fmt.Println("---------------------")
	fmt.Println("Enter the command")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\r\n", "", -1)

		if strings.Compare("help", text) == 0 {
			fmt.Println("You can use the next commands:")
			fmt.Println("pin: mark a file as important (can't be removed)")
			fmt.Println("unpin: dismark the file as important (can be removed)")
			fmt.Println("cat: print the content of the file given")
			fmt.Println("store: save a file in one of the nodes")
			fmt.Println("exit: leave the simulation")
		} else if strings.Compare("cat", text) == 0 {
			fmt.Println("Debe de leer el dato almacenado")
		} else if strings.Compare("pin", text) == 0 {
			fmt.Println("Bloquea un dato para que no pueda ser borrado")
		} else if strings.Compare("unpin", text) == 0 {
			fmt.Println("Desbloquea un dato para poder mamonearlo como quieras")
		} else if strings.Compare("store", text) == 0 {
			fmt.Println("Almacena un dato")
		} else if strings.Compare("exit", text) == 0 {
			break
		} else {
			fmt.Println("Wrong command. Please try again")
			fmt.Println("If you need any help, use the help command")
		}

	}

}

/*
// Package implementing formatted I/O.

type node struct {
	ip        string
	port      int
	networkID int
}

func main() {
	srcContact := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001")
	rt := NewRoutingTable(srcContact)

	otherContact := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000")
	rt.AddContact(otherContact)
	rt.AddContact(NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002"))

	contacts := rt.FindClosestContacts(NewKademliaID("2111111400000000000000000000000000000000"), 20)
	for i := range contacts {
		fmt.Println(contacts[i].String())
	}

	kademlia := NewKademlia(*rt, 2, 1)
	go kademlia.ServerThread("8001")
	fmt.Println(SendPingMessage(srcContact, otherContact))
}
*/
