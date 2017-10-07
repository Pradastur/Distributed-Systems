package main

import "time"

type File struct {
	path           string
	pinned         bool
	expirationDate time.Time
	content        []byte
}

func NewFile(pathA string, pinn bool, cont []byte) File {
	return File{pathA, pinn, time.Now().Add(time.Hour * 24), cont}
}

func (file *File) IsPinned() bool {
	return file.pinned
}

func (file *File) IsOutDate() bool {
	return !file.expirationDate.After(time.Now())
}
