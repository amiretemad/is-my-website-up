package main

import (
	"main/Interfaces"
	"main/MessageChannels"
	"main/Structs"
)

func MessageSender(l *Structs.LinkStruct)  {

	a := []Interfaces.SendMessage{
		&MessageChannels.SlackChannel{L: *l},
		&MessageChannels.MatterMostChannel{L: *l},
	}

	for _, messagingChannels := range a {
		messagingChannels.SendMessage()
	}

}
