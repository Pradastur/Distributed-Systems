package main

import "fmt"

type savedata struct {
	dataSaved []byte
}

func (data *savedata) Save(saving []byte) {
	data.dataSaved = saving
	fmt.Println("File saved")
}

func (data *savedata) Get() []byte {
	return data.dataSaved
}
