package main

import (
	"encoding/xml"
	"io/ioutil"
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

func NewFileSystem(path string) Filesystem {
	xmlData := ReadFile(path)
	var XMLF XMLCreator
	xml.Unmarshal(xmlData, &XMLF)

	files := make(map[string]File)
	for i := range XMLF.FileArray {
		file := XMLF.FileArray[i]

		loadF := LoadFile(file.Path, file.Pinned, file.ExpirationDate)
		files[Hash(loadF.Path)] = loadF
	}
	return Filesystem{NewFile(path, true, xmlData), files}
}

func ReadFile(path string) []byte {
	file, _ := os.Open(path)
	byteString, _ := ioutil.ReadAll(file)
	file.Close()

	return byteString
}

func (fileSystem *Filesystem) UpdateFile() {
	fileList := make([]File, len(fileSystem.files))
	i := 0
	for _, file := range fileSystem.files {
		fileList[i] = file
		i = i + 1
	}

	creator := XMLCreator{FileArray: fileList}
	newFile, _ := os.Create(fileSystem.configFile.Path)
	contentFile, _ := xml.MarshalIndent(creator, " ", " 	")
	newFile.Write(contentFile)
	newFile.Close()

}

func (fileSystem *Filesystem) save(file File) {
	hash := Hash(file.Path)
	if file.Path != fileSystem.configFile.Path {
		if !fileSystem.hasData(hash) {
			fileSystem.files[hash] = file
			newFile, _ := os.Create("data/" + file.Path)
			newFile.Write(file.Content)
			newFile.Close()
			fileSystem.UpdateFile()
			//TODO Falta XML

		}
	}
}

func (fileSystem *Filesystem) remove(file File) {
	hash := Hash(file.Path)
	if fileSystem.hasData(hash) {
		if !file.IsPinned() {
			fileSystem.files[hash] = File{}
			os.Remove(hash)
			//TODO Falta XML
			fileSystem.UpdateFile()
		}
	}
}
