package Structs

type LinkStruct struct {
	Label       string
	Link        string
	Description string
	IsUp        bool
}

func NewLinkStruct() LinkStruct {
	return LinkStruct{}
}

func (l *LinkStruct) GetLinks() []LinkStruct {
	return []LinkStruct{
		{
			Label:       "Amazon",
			Link:        "https://amasdzon.com",
			Description: "Amazon Website",
			IsUp:        false,
		},
		{
			Label:       "Spotify",
			Link:        "http://spoaqwtify.com",
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
}
