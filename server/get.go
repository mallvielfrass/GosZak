package main

import (
	//"fmt"
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
	OneBlock := ""
	massivData := ""
	//fmt.Println(doc.Html())
	header, _ := doc.Find(".allRecords").Html()
	//content, _ := doc.Find(".parametrs").Html()
	//fmt.Println(header)\
	doc.Find("div").Each(func(i int, s *goquery.Selection) {
		class, _ := s.Attr("class")
		switch class {
		case "registerBox registerBoxBank margBtm20":
			OneBlock, _ = s.Html()
			massivData = massivData + OneBlock
			//fmt.Println(s.Html())
		default:

		}
	})

	//fmt.Println(content, doc.Find(".allRecords").Text())
	returned := massivData + string(header)
	return returned
}
