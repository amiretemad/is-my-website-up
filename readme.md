# Is my website up ?

A simple script to check your website is up or down and you will be notified from **Slack** and **Mattermost**

![Image](https://i.ibb.co/Sm1mXJX/Screen-Shot-2020-08-06-at-2-36-54-PM.png)

## Description

A simple script to check your website is up or down by adding lists of websites to Links Struct, you will able to check the status of your website.


## Supported Messaging channels 

 - Slack (ready to use)
 - Mattermost (ready to use)

## Executing program

Running Script : 

```
go run main.go
```

Building Script: 

```
go build main.go
```

## Define your links inside Structs/LinkStruct.go

You can easily add your own links inside this list of links.
```
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
```

## Creating .env file
make a copy of .env.example in root directory of project as .env and enter your Slack Hook URL

```
SLACK_HOOK_URL=YOUR SLACK HOOK URL
MATTERMOST_HOOK_URL=YOUR MATTERMOST HOOK URL
```
