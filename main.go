package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//	file := NewFile("/data", true, []byte(" "))
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("How many nodes to deploy")

	nNode, _ := reader.ReadString('\n')
	nNode = strings.Replace(nNode, "\r\n", "", -1)
	nodeInt, _ := strconv.Atoi(nNode)

	for i := 0; i < nodeInt; i++ {
		if i < 10 {
			port := i
			KademliaIDName := NewKademliaID("000000000000000000000000000000000000000" + strconv.Itoa(port))
			portNext := i + 1
			if i != 9 {
				KademliaIDNameNext := NewKademliaID("000000000000000000000000000000000000000" + strconv.Itoa(portNext))
				go nextNode(KademliaIDName, KademliaIDNameNext, port, portNext)
			} else {
				KademliaIDNameNext2 := NewKademliaID("00000000000000000000000000000000000000" + strconv.Itoa(portNext))
				go nextNode(KademliaIDName, KademliaIDNameNext2, port, portNext)
			}
			//fmt.Println(KademliaIDName.String())
			//KademliaIDnode++
		} else {
			port := i
			KademliaIDName := NewKademliaID("00000000000000000000000000000000000000" + strconv.Itoa(port))
			portNext := i + 1
			KademliaIDNameNext := NewKademliaID("00000000000000000000000000000000000000" + strconv.Itoa(portNext))

			//fmt.Println(KademliaIDName)
			//KademliaIDnode++
			go nextNode(KademliaIDName, KademliaIDNameNext, port, portNext)
		}
	}
	finalNode(nodeInt)
	/*
		mySelf := NewContact(NewKademliaID(firstAddress), "localhost: "+strconv.Itoa(firstPortInt))
		rt := NewRoutingTable(mySelf)
		channel := make(chan []Contact)
		kademlia := NewKademlia(*rt, 20, 3, channel)

		go kademlia.network.Listen("localhost", firstPortInt)
	*/

}

func nextNode(kademliaIDName *KademliaID, kademliaIDNameNext *KademliaID, port int, portNext int) {
	mySelf := NewContact(kademliaIDName, "localhost:"+strconv.Itoa(8000+port))
	next := NewContact(kademliaIDNameNext, "localhost:"+strconv.Itoa(8000+portNext))

	rt := NewRoutingTable(mySelf)

	rt.AddContact(mySelf)
	rt.AddContact(next)

	channel := make(chan []Contact)
	kademlia := NewKademlia(*rt, 20, 3, channel)
	file := NewFile("/data/"+strconv.Itoa(port)+".txt", true, []byte(strconv.Itoa(port)))
	kademlia.Store(file)
	//	fileSystem := kademlia.fSys

	go kademlia.network.Listen("localhost", 8000+port)
}

func finalNode(nodeInt int) {
	var KademliaIDName *KademliaID
	if nodeInt < 10 {
		KademliaIDName = NewKademliaID("000000000000000000000000000000000000000" + strconv.Itoa(nodeInt))
		/*
			if nodeInt != 9 {
				KademliaIDName = NewKademliaID("000000000000000000000000000000000000000" + strconv.Itoa(nodeInt))
			} else {
				KademliaIDName = NewKademliaID("00000000000000000000000000000000000000" + strconv.Itoa(nodeInt))
			}
		*/
		//fmt.Println(KademliaIDName.String())
		//KademliaIDnode++
	} else {
		KademliaIDName = NewKademliaID("00000000000000000000000000000000000000" + strconv.Itoa(nodeInt))
	}

	mySelf := NewContact(KademliaIDName, "localhost:"+strconv.Itoa(8000+nodeInt))
	node0 := NewContact(NewKademliaID("0000000000000000000000000000000000000000"), "localhost: 8000")

	rt := NewRoutingTable(mySelf)

	rt.AddContact(mySelf)
	rt.AddContact(node0)

	channel := make(chan []Contact)
	kademlia := NewKademlia(*rt, 20, 3, channel)
	fileSystem := kademlia.fSys

	go kademlia.network.Listen("localhost", 8000+nodeInt)

	//file := NewFile("/data", true, []byte(" "))
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("")
	fmt.Println("WELCOME TO KADEMLIA")
	fmt.Println("---------------------")
	fmt.Println("Enter the command")
	fileList := make([]File, len(fileSystem.files))

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\r\n", "", -1)

		if strings.Compare("help", text) == 0 {
			fmt.Println("You can use the next commands:")
			//fmt.Println("new [int nodes]: create a new scenario. You have to indicate the number of nodes")
			fmt.Println("newFile [string content]: create a new file pinned")
			fmt.Println("join [contact Contact]: join to the network using the contact passed as parameter")
			fmt.Println("pin [file File]: mark a file as important (can't be removed)")
			fmt.Println("remove [file File]: remove this file if its not pinned")
			fmt.Println("unpin [file File]: dismark the file as important (can be removed)")
			fmt.Println("cat [file File]: print the content of the file given")
			fmt.Println("store [file File] [string address]: save a file in the node")
			fmt.Println("exit: leave the simulation")
			/*
				fmt.Println("Introduce a valid address of 40 numbers, example --> [0000000000000000000000000000000000000000]")
				address, _ := reader.ReadString('\n')
				address = strings.Replace(address, "\r\n", "", -1)

				fmt.Println("Address introduced: " + address)

				fmt.Println("Introduce a valid port of 4 numbers, example --> 8000")
				port, _ := reader.ReadString('\n')
				port = strings.Replace(port, "\r\n", "", -1)

				fmt.Println("Port introduced: " + port)
				portInt, _ := strconv.Atoi(port)

				go node(address, portInt)
			*/
		} else if strings.Compare("join", text) == 0 {
			fmt.Println("Une el nodo en la red usando el contacto pasado por parametro")
			//var contact Contact
			//Join(contact)
		} else if strings.Compare("cat", text) == 0 {

			fmt.Println("Select the file to see the content")
			//fileList := make([]File, len(fileSystem.files))
			i := 0
			for _, file := range fileSystem.files {
				fileList[i] = file
				fmt.Println(i, ": ", file.Path)
				fmt.Println(i, ": ", file.Content)
				i = i + 1
			}
			fmt.Println("Select which one you want to see, introducing the number associated")
			numFile, _ := reader.ReadString('\n')
			numFile = strings.Replace(numFile, "\r\n", "", -1)
			//	fileWanted, _ := strconv.Atoi(numFile)
			fmt.Println("Este es: " + numFile)
			numFileInt, _ := strconv.Atoi("" + numFile)
			//no  hace el cat del file porque no ve que lo que introducimos es un numero
			fileW := fileList[numFileInt]
			content := string(fileW.Content)
			fmt.Println("Content is: ", content)

		} else if strings.Compare("pin", text) == 0 {

			fmt.Println("Select the file to pin")
			//fileList := make([]File, len(fileSystem.files))
			i := 0
			for _, file := range fileSystem.files {
				fileList[i] = file
				fmt.Println(i, ": ", file.Path)
				i = i + 1
			}
			fmt.Println("Select which one you want to pin, introducing the number associated")
			numFile, _ := reader.ReadString('\n')
			numFile = strings.Replace(numFile, "\r\n", "", -1)
			fileWanted, _ := strconv.Atoi(numFile)
			fileW := fileList[fileWanted]
			fileW.Pinned = true
			fmt.Println("File pinned")

		} else if strings.Compare("remove", text) == 0 {

			fmt.Println("Select the file to remove")
			//fileList := make([]File, len(fileSystem.files))
			i := 0
			for _, file := range fileSystem.files {
				fileList[i] = file
				fmt.Println(i, ": ", file.Path)
				fmt.Println(i, ": ", file.IsPinned())
				i = i + 1
			}
			fmt.Println("Select which one you want to remove, introducing the number associated")
			numFile, _ := reader.ReadString('\n')
			numFile = strings.Replace(numFile, "\r\n", "", -1)
			fileWanted, _ := strconv.Atoi(numFile)
			fileW := fileList[fileWanted]
			/*if !fileW.IsPinned() {
				fileSystem.remove(fileW)
				fmt.Println("File removed")
			} else {
				fmt.Println("Unremovable")
			}*/
			fileSystem.remove(fileW)

		} else if strings.Compare("unpin", text) == 0 {

			fmt.Println("Select the file to unpin")
			//fileList := make([]File, len(fileSystem.files))
			i := 0
			for _, file := range fileSystem.files {
				fileList[i] = file
				fmt.Println(i, ": ", file.Path)
				i = i + 1
			}
			fmt.Println("Select which one you want to unpin, introducing the number associated")
			numFile, _ := reader.ReadString('\n')
			numFile = strings.Replace(numFile, "\r\n", "", -1)
			fileWanted, _ := strconv.Atoi(numFile)
			fileW := fileList[fileWanted]
			fileW.Pinned = false
			fmt.Println("File unpinned")

		} else if strings.Compare("newFile", text) == 0 {
			fmt.Println("Introduce the content of the file")
			cont, _ := reader.ReadString('\n')
			cont = strings.Replace(cont, "\r\n", "", -1)

			fmt.Println("Name of the file and type, ej: file.txt")
			nameFile, _ := reader.ReadString('\n')
			nameFile = strings.Replace(nameFile, "\r\n", "", -1)

			file := NewFile(nameFile, true, []byte(cont))

			fileSystem.save(file)

			fmt.Println("*****File Created*****")
		} else if strings.Compare("store", text) == 0 {

			fmt.Println("Select the file to store")
			//fileList := make([]File, len(fileSystem.files))
			i := 0
			for _, file := range fileSystem.files {
				fileList[i] = file
				fmt.Println(i, ": ", file.Path)
				i = i + 1
			}
			fmt.Println("Select which one you want to store, introducing the number associated")
			numFile, _ := reader.ReadString('\n')
			numFile = strings.Replace(numFile, "\r\n", "", -1)
			fileWanted, _ := strconv.Atoi(numFile)
			fileW := fileList[fileWanted]
			kademlia.Store(fileW)

			fmt.Println("File stored")
		} else if strings.Compare("exit", text) == 0 {
			break
		} else {
			fmt.Println("Wrong command. Please try again")
			fmt.Println("If you need any help, use the help command")
		}

	}

}

/*func node(address string, port int) {
	mySelf := NewContact(NewKademliaID(address), "localhost: "+strconv.Itoa(port))

	rt := NewRoutingTable(mySelf)

	rt.AddContact(mySelf)

	channel := make(chan []Contact)
	kademlia := NewKademlia(*rt, 20, 3, channel)
	fileSystem := kademlia.fSys

	go kademlia.network.Listen("localhost", port)

}*/

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
