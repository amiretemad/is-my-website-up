package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"main/Structs"
	"net/http"
	"time"
)

const colorRed = "\033[31m"
const colorGreen = "\033[32m"

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c := make(chan Structs.LinkStruct)

	newLinkStruct := Structs.NewLinkStruct()
	for _, l := range newLinkStruct.GetLinks() {
		go checkLinkHealth(l, c)
	}

	for l := range c {
		if l.IsUp {
			fmt.Println(colorGreen, l.Label+" Is Up ("+l.Link+")")
		} else {
			fmt.Println(colorRed, l.Label+" Is Down ("+l.Link+")")
			MessageSender(&l)
		}
		time.Sleep(time.Second * 2)
		go checkLinkHealth(l, c)
	}
}

func checkLinkHealth(l Structs.LinkStruct, c chan Structs.LinkStruct) {
	_, err := http.Get(l.Link)

	if err != nil {
		l.IsUp = false
		c <- l
		return
	}

	l.IsUp = true
	c <- l
}
