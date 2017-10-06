package main

type Filesystem struct {
	files []FileKademlia
}

type FileKademlia struct {
	path           string
	pinned         bool
	expirationDate time
}

func hasData(hash string) bool {
	return true
}

func get(hash string) Filesystem {

}

func save(file FileKademlia) {

}

func remove(hash string) {

}
