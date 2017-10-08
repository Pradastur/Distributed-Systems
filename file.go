package main

import "time"

type File struct {
	Path           string    `xml: "Path"`
	Pinned         bool      `xml: "isPinned"`
	ExpirationDate time.Time `xml: "UntilDate"`
	Content        []byte    `xml: "-"`
}

func NewFile(PathA string, pinn bool, cont []byte) File {
	return File{PathA, pinn, time.Now().Add(time.Hour * 24), cont}
}

func (file *File) IsPinned() bool {
	return file.Pinned
}

func (file *File) IsOutDate() bool {
	return !file.ExpirationDate.After(time.Now())
}
