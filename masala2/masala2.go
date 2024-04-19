package main

import (
	"fmt"
	"log"
	"os"
)

func FilePlusCopy(){
	malumot,err := os.ReadFile("file1.txt")
	if err != nil{
		log.Fatal(err)
	}

	malumot1,err := os.ReadFile("file2.txt")
	if err != nil{
		log.Fatal(err)
	} 

	file,err := os.OpenFile("file3.txt",os.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_TRUNC,0666)
	if err != nil{
		log.Fatal(err)
	}
	file.Write(malumot)
	file.Write(malumot1)
	var malumot3 string
	malumot3=string(malumot)+" "+string(malumot1)
	fmt.Println(malumot3)
}

func main(){
	FilePlusCopy()
}