package main

import (
	"os"
)

type Filesystem struct {
	configFile File
	files      map[string]File
}

func (fileSystem *Filesystem) hasData(hash string) bool {
	return fileSystem.files[hash].path != ""
}

func (fileSystem *Filesystem) getFile(hash string) File {
	if fileSystem.hasData(hash) {
		return fileSystem.files[hash]
	}
	return File{}
}

func (fileSystem *Filesystem) save(file File) {
	hash := Hash(file.content)
	if file.path != fileSystem.configFile.path {
		if fileSystem.hasData(hash) {
			fileSystem.files[hash] = file

			newFile, _ := os.Create("data/" + file.path)
			newFile.Write(file.content)
			newFile.Close()
			//TODO Falta XML
		}
	}
}

func (fileSystem *Filesystem) remove(file File) {
	hash := Hash(file.content)
	if fileSystem.hasData(hash) {
		if !file.IsPinned() {
			fileSystem.files[hash] = File{}
			os.Remove(hash)
			//TODO Falta XML
		}
	}
}
