package main

import (
	"log"

	"github.com/PuerkitoBio/goquery"
)

func get(params string) string {
	url := "http://www.zakupki.gov.ru/epz/order/quicksearch/search.html?" + params
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	return string(b)
}
