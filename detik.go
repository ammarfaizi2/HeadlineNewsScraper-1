package main

import (
	"fmt"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
)

func getScrapeDetik() {
	site := "http://detik.com"
	doc, err := goquery.NewDocument(site)
	if err != nil {
		log.Fatal(err)
	}

	//headline
	fmt.Printf("headline\n")
	//bertiaiutama
	doc.Find(".beritautama .img_con").Each(func(index int, item *goquery.Selection) {
		title := rws(item.Text())
		insertDB(sha1gen(title), site, title, time.Now())
	})
	fmt.Printf("\n")

	//landscape
	doc.Find(".box_hl_new a").EachWithBreak(func(index int, item *goquery.Selection) bool {
		title := rws(item.Text())
		if title == "" {
			return true
		}
		insertDB(sha1gen(title), site, title, time.Now())
		return true
	})
	fmt.Printf("\n")

	//r_content
	fmt.Printf("r_content\n")
	//h3
	doc.Find(".r_content h3").Each(func(index int, item *goquery.Selection) {
		title := rws(item.Text())
		insertDB(sha1gen(title), site, title, time.Now())
	})
	fmt.Printf("\n")

	//h2
	doc.Find(".r_content h2").Each(func(index int, item *goquery.Selection) {
		title := rws(item.Text())
		insertDB(sha1gen(title), site, title, time.Now())
	})
	fmt.Printf("\n")

	//l_content
	fmt.Printf("l_content\n")
	doc.Find(".l_content h2").Each(func(index int, item *goquery.Selection) {
		title := rws(item.Text())
		insertDB(sha1gen(title), site, title, time.Now())
	})
	fmt.Printf("\n")

	//newsfeed
	fmt.Printf("newsfeed\n")
	doc.Find("#newsfeed-anchor-container h2").Each(func(index int, item *goquery.Selection) {
		title := rws(item.Text())
		insertDB(sha1gen(title), site, title, time.Now())
	})
	fmt.Printf("\n")
}
