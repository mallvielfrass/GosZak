package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func get(params string) string {
	url := "http://www.zakupki.gov.ru/epz/order/quicksearch/search.html?" + params
	page, err := http.Get(url) // выполняем запрос
	if err != nil {
		log.Fatalln(err)
	}
	defer page.Body.Close()
	if page.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", page.StatusCode, page.Status)
	}

	doc, err := goquery.NewDocumentFromReader(page.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(doc.Html())
	header := doc.Find(".allRecords").Text()
	content, _ := doc.Find(".parametrs").Html()
	//fmt.Println(header)

	fmt.Println(content, doc.Find(".allRecords").Text())
	return string(header + content)
}
