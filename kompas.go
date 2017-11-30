package main

import (
	"fmt"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
)

func getScrapeKompas() {
	site := "http://kompas.com"
	doc, err := goquery.NewDocument(site)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".headline__big__title").Each(func(index int, item *goquery.Selection) {
		title := rws(item.Text())
		insertDB(sha1gen(title), site, title, time.Now())
	})
	fmt.Printf("\n")

	doc.Find(".most__title").Each(func(index int, item *goquery.Selection) {
		title := rws(item.Text())
		insertDB(sha1gen(title), site, title, time.Now())
	})
	fmt.Printf("\n")

	doc.Find(".article__title").Each(func(index int, item *goquery.Selection) {
		title := rws(item.Text())
		insertDB(sha1gen(title), site, title, time.Now())
	})
	fmt.Printf("\n")

	doc.Find(".opinion__title").Each(func(index int, item *goquery.Selection) {
		title := rws(item.Text())
		insertDB(sha1gen(title), site, title, time.Now())
	})
}
