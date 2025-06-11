package handlers

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"strings"
	"vknow360/ie"
	"vknow360/toi"

	"github.com/gin-gonic/gin"
)

type NewsItem struct {
	Title      string   `json:"title"`
	Link       string   `json:"link"`
	PubDate    string   `json:"pubDate"`
	Categories []string `json:"categories"`
	Enclosure  string   `json:"enclosure"`
	Source     string   `json:"source"`
}

func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Status": "Okay"})
}

func GetNews(c *gin.Context) {
	urls := []string{
		"https://timesofindia.indiatimes.com/rssfeeds/2647163.cms",
		"https://indianexpress.com/section/weather/feed",
		"https://indianexpress.com/section/climate-change/feed"}
	var res *http.Response
	var err error

	news := []NewsItem{}

	for index, url := range urls {
		res, err = http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		body, _ := io.ReadAll(res.Body)

		if index == 0 {
			var toi toi.TimesOfIndia

			err = xml.Unmarshal(body, &toi)
			if err != nil {
				log.Fatal(err)
			}
			for _, item := range toi.Channel.Items {
				news = append(news, NewsItem{
					Title:     item.Title,
					Link:      item.Link,
					PubDate:   item.PubDate,
					Enclosure: item.Enclosure.Url,
					Categories: func() []string {
						cats := []string{"weather"}
						if strings.Contains(item.Title, "climate change") {
							cats = append(cats, "climate")
						}
						return cats
					}(),
					Source: "Times of India",
				})
			}
		} else {
			var ie ie.IndianExpress

			err = xml.Unmarshal(body, &ie)
			if err != nil {
				log.Fatal(err)
			}
			categories := []string{}
			if index == 1 {
				categories = append(categories, "weather")
			} else {
				categories = append(categories, "climate")
			}
			for _, item := range ie.Channel.Items {
				news = append(news, NewsItem{
					Title:      item.Title,
					Link:       item.Link,
					PubDate:    item.PubDate,
					Categories: categories,
					Enclosure:  item.MediaContent.URL,
					Source:     "Indian Express",
				})
			}
		}

	}

	c.JSON(http.StatusOK, gin.H{"news": news})
}
