package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

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
		} else {
			port := i
			KademliaIDName := NewKademliaID("00000000000000000000000000000000000000" + strconv.Itoa(port))
			portNext := i + 1
			KademliaIDNameNext := NewKademliaID("00000000000000000000000000000000000000" + strconv.Itoa(portNext))

			go nextNode(KademliaIDName, KademliaIDNameNext, port, portNext)
		}
	}
	finalNode(nodeInt)
}

func nextNode(kademliaIDName *KademliaID, kademliaIDNameNext *KademliaID, port int, portNext int) {
	mySelf := NewContact(kademliaIDName, "localhost:"+strconv.Itoa(8000+port))
	next := NewContact(kademliaIDNameNext, "localhost:"+strconv.Itoa(8000+portNext))

	rt := NewRoutingTable(mySelf)

	rt.AddContact(mySelf)
	rt.AddContact(next)

	channel := make(chan []Contact)
	kademlia := NewKademlia(*rt, 20, 3, channel)
	go kademlia.network.Listen("localhost", 8000+port)
}

func finalNode(nodeInt int) {
	var KademliaIDName *KademliaID
	if nodeInt < 10 {
		KademliaIDName = NewKademliaID("000000000000000000000000000000000000000" + strconv.Itoa(nodeInt))
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

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("")
	fmt.Println("WELCOME TO KADEMLIA")
	fmt.Println("---------------------")
	fmt.Println("Enter the command")

	for {
		fileList := make([]File, len(fileSystem.files))
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\r\n", "", -1)

		if strings.Compare("help", text) == 0 {
			fmt.Println("You can use the next commands:")
			fmt.Println("newFile [string content]: create a new file pinned")
			fmt.Println("pin [file File]: mark a file as important (can't be removed)")
			fmt.Println("remove [file File]: remove this file if its not pinned")
			fmt.Println("unpin [file File]: dismark the file as important (can be removed)")
			fmt.Println("cat [file File]: print the content of the file given")
			fmt.Println("store [file File] [string address]: save a file in the node")
			fmt.Println("exit: leave the simulation")
		} else if strings.Compare("cat", text) == 0 {
			fileSystem.UpdateFile()

			fmt.Println("Select the file to see the content")
			i := 0
			for _, file := range fileSystem.files {
				fileList[i] = file
				fmt.Println(i, ": ", file.Path)
				i = i + 1
			}

			fmt.Println("Select which one you want to see, introducing the number associated")
			numFile, _ := reader.ReadString('\n')
			numFile = strings.Replace(numFile, "\r\n", "", -1)
			fileWanted, _ := strconv.Atoi(numFile)
			fileW := fileList[fileWanted]
			content := string(fileW.Content)
			fmt.Println("Content is: ", content)

		} else if strings.Compare("pin", text) == 0 {
			fileSystem.UpdateFile()

			fmt.Println("Select the file to pin")
			i := 0
			for _, file := range fileSystem.files {
				fileList[i] = file
				fmt.Println(i, ": ", file.Path)
				fmt.Println(i, ": ", file.IsPinned())
				i = i + 1
			}
			fmt.Println("Select which one you want to pin, introducing the number associated")
			numFile, _ := reader.ReadString('\n')
			numFile = strings.Replace(numFile, "\r\n", "", -1)
			fileWanted, _ := strconv.Atoi(numFile)
			fileW := fileList[fileWanted]
			fileSystem.PinFile(fileW)
			fmt.Println("File pinned")

		} else if strings.Compare("remove", text) == 0 {

			fmt.Println("Select the file to remove")
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
			if !fileW.IsPinned() {
				fileSystem.remove(fileW)
				fmt.Println("File removed")
			} else {
				fmt.Println("Unremovable")
			}

		} else if strings.Compare("unpin", text) == 0 {
			fileSystem.UpdateFile()

			fmt.Println("Select the file to unpin")
			i := 0
			for _, file := range fileSystem.files {
				fileList[i] = file
				fmt.Println(i, ": ", file.Path)
				fmt.Println(i, ": ", file.IsPinned())
				i = i + 1
			}
			fmt.Println("Select which one you want to unpin, introducing the number associated")
			numFile, _ := reader.ReadString('\n')
			numFile = strings.Replace(numFile, "\r\n", "", -1)
			fileWanted, _ := strconv.Atoi(numFile)
			fileW := fileList[fileWanted]
			fileSystem.UnpinFile(fileW)
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
			fileSystem.UpdateFile()

		} else if strings.Compare("store", text) == 0 {

			fmt.Println("Select the file to store")
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
			fileSystem.UpdateFile()
		} else if strings.Compare("exit", text) == 0 {
			break
		} else {
			fmt.Println("Wrong command. Please try again")
			fmt.Println("If you need any help, use the help command")
		}

	}

}
