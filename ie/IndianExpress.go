package ie

type IndianExpress struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Items []Item `xml:"item"`
}

type Item struct {
	Title        string       `xml:"title"`
	Link         string       `xml:"link"`
	PubDate      string       `xml:"pubDate"`
	MediaContent MediaContent `xml:"content"`
}

type MediaContent struct {
	URL string `xml:"url,attr"`
}
