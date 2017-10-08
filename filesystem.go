package main

import (
	"encoding/xml"
	"os"
)

type Filesystem struct {
	configFile File
	files      map[string]File
}

type XMLCreator struct {
	XmlName   xml.Name `xml:"files"`
	FileArray []File   `xml:"file"`
}

func (fileSystem *Filesystem) hasData(hash string) bool {
	return fileSystem.files[hash].Path != ""
}

func (fileSystem *Filesystem) getFile(hash string) File {
	if fileSystem.hasData(hash) {
		return fileSystem.files[hash]
	}
	return File{}
}

func (fileSystem *Filesystem) UpdateFile() {
	fileList := make([]File, len(fileSystem.files))
	i := 0
	for hash, file := range fileSystem.files {
		fileList[i] = file
		i = i + 1
	}

	creator := XMLCreator{FileArray: fileList}
	newFile, _ := os.Create("data/" + fileSystem.configFile.Path)
	contentFile, _ := xml.MarshalIndent(creator, " ", " 	")
	newFile.Write(contentFile)
	newFile.Close()

}

func (fileSystem *Filesystem) save(file File) {
	hash := Hash(file.Content)
	if file.Path != fileSystem.configFile.Path {
		if fileSystem.hasData(hash) {
			fileSystem.files[hash] = file

			newFile, _ := os.Create("data/" + file.Path)
			newFile.Write(file.Content)
			newFile.Close()
			//TODO Falta XML

		}
	}
}

func (fileSystem *Filesystem) remove(file File) {
	hash := Hash(file.Content)
	if fileSystem.hasData(hash) {
		if !file.IsPinned() {
			fileSystem.files[hash] = File{}
			os.Remove(hash)
			//TODO Falta XML
		}
	}
}
