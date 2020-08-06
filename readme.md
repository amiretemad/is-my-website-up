# Is my website up ?

A simple script to check your website is up or down.

## Description

A simple script to check your website is up or down by adding list of websites to Links Struct inside main.go you are able to check the status of your website 

## Executing program

Running Script : 

```
go run main.go
```

Building Script: 

```
go build main.go
```

## Define your links inside linkStruct

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
