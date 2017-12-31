package main

import "testing"
import "fmt"
import "github.com/PuerkitoBio/goquery"

func TestProfile(t *testing.T) {
	doc, err := goquery.NewDocument("https://twitter.com/konojunya")
	if err != nil {
		t.Error(err)
	}

	name := doc.Find(".ProfileHeaderCard-nameLink").Text()
	icon, exists := doc.Find(".ProfileAvatar-image").Attr("src")
	if exists {
		fmt.Println(name, icon)
	}

}
