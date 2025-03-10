package models

import (
	"fmt"
	"time"
)

type Category struct {
	Id   int64  `json:"id"`
	Tag  string `json:"tag"`
	Name string `json:"name"`
}

type Post struct {
	Id               int64  `json:"id"`
	Text             string `json:"text"`
	Ref              string `json:"ref"`
	PublicationDate  string `json:"publication_date"`
	LogoRef          string `json:"logo_ref"`
	ChannelName      string `json:"channel_name"`
	SubscribersCount int64  `json:"subscribers_count"`
	CategoryId       int64  `json:"category_id"`
	Views            int64  `json:"views"`
	Redirected       int    `json:"redirected"`
	IsChannelPremium bool   `json:"premium"`
}

func (p Post) GetRating() float64 {
	publicationDate, err := time.Parse("2006-01-02", p.PublicationDate)
	if err != nil {
		fmt.Println("Problem with parsing date")
		return 0.0
	}
	currentDate := time.Now().Local()
	days := currentDate.Sub(publicationDate).Hours() / 24
	return float64(p.SubscribersCount)*0.001 + float64(p.Redirected)/float64(p.Views)*1000 - days*100
}
