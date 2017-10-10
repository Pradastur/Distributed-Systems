package main

import (
	"time"
)

type File struct {
	Path           string    `xml: "Path"`
	Pinned         bool      `xml: "isPinned"`
	ExpirationDate time.Time `xml: "UntilDate"`
	Content        []byte    `xml: "-"`
}

func NewFile(PathA string, pinn bool, cont []byte) File {
	return File{PathA, pinn, time.Now().Add(time.Hour * 24), cont}
}

func LoadFile(PathA string, pinn bool, expDate time.Time, cont []byte) File {
	//content := ReadFile("/data/" + PathA)
	return File{PathA, pinn, expDate, cont}
}

func (file *File) IsPinned() bool {
	return file.Pinned
}

func (file *File) IsOutDate() bool {
	return !file.ExpirationDate.After(time.Now())
}
