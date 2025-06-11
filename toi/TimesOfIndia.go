package toi

type TimesOfIndia struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Items []Item `xml:"item"`
}

type Item struct {
	Title     string    `xml:"title"`
	Link      string    `xml:"link"`
	PubDate   string    `xml:"pubDate"`
	Enclosure Enclosure `xml:"enclosure"`
}

type Enclosure struct {
	Type string `xml:"type,attr"`
	Url  string `xml:"url,attr"`
}
