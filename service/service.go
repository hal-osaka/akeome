package service

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/alberii/akeome/model"
)

var db = model.GetDBConn()

func GetProfile(id string) (icon string, nickname string, err error) {
	doc, err := goquery.NewDocument("https://twitter.com/" + id)
	if err != nil {
		return "", "", err
	}

	name := doc.Find(".ProfileHeaderCard-nameLink").Text()
	icon, exists := doc.Find(".ProfileAvatar-image").Attr("src")
	if exists {
		fmt.Println(name, icon)
	}

	return icon, name, nil
}
