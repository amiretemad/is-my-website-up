package main

import (
	"fmt"
	"net/http"
	"time"
)

type linkStruct struct {
	Label       string
	Link        string
	Description string
	IsUp        bool
}

const colorRed = "\033[31m"
const colorGreen = "\033[32m"

func main() {

	links := []linkStruct{
		{
			Label:       "Amazon",
			Link:        "https://amazon.com",
			Description: "Amazon Website",
			IsUp:        false,
		},
		{
			Label:       "Spotify",
			Link:        "http://spotify.com",
			Description: "Spotify Website",
			IsUp:        false,
		},
		{
			Label:       "Google",
			Link:        "https://google.com",
			Description: "Google Main Website",
			IsUp:        false,
		},
	}

	// Making Channel
	c := make(chan linkStruct)

	for _, l := range links {
		go checkLinkHealth(l, c)
	}

	for l := range c {
		time.Sleep(time.Second * 2)
		go checkLinkHealth(l, c)
		if l.IsUp {
			fmt.Println(colorGreen, l.Label+" Is Up ("+l.Link+")")
		} else {
			fmt.Println(colorRed, l.Label+" Is Down ("+l.Link+")")
		}
	}

}

func checkLinkHealth(l linkStruct, c chan linkStruct) {
	_, err := http.Get(l.Link)

	if err != nil {
		l.IsUp = false
		c <- l
		return
	}

	l.IsUp = true
	c <- l
}
