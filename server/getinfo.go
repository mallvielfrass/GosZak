package main

import (
	"fmt"
	"log"
	"time"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type GetInfo struct {}

type StructOneBlock struct {
	BoxIcons         string
	TenderTd         string
	DescriptTenderTd string
	AmountTenderTd   string
	ReportBox        string
}

func (_ GetInfo) SearchQueryToParams(searchQuery SearchQuery) string {
	params := ""

	for _, item := range searchQuery.LawNumber {
		if len(params) > 0 {
			params += "&"
		}

		if item == "94-fz" {
			params += "fz94=on"
		}

		if item == "pp_rf_615" {
			params += "ppRf615=on"
		}

		if item == "223-fz" {
			params += "fz223=on"
		}

		if item == "44-fz" {
			params += "fz44=on"
		}
	}

	for _, item := range searchQuery.ProcedureType {
		if len(params) > 0 {
			params += "&"
		}

		if item == "applicationSubmission" {
			params += "af=on"
		}

		if item == "commissionWork" {
			params += "ca=on"
		}

		if item == "procedureCompleted" {
			params += "pc=on"
		}

		if item == "procedureAborted" {
			params += "pa=on"
		}
	}

	if len(searchQuery.CityName) > 0 {
		if len(params) > 0 {
			params += "&"
		}

		switch searchQuery.CityName {
		case "st_petersburg":
			params += "selectCity=1&districts=5277336&region_regions_5277347=region_regions_5277347&regions=5277347"
		default:
			params += "selectCity=" + url.QueryEscape(searchQuery.CityName)
		}
	}

	if searchQuery.PublishDateFrom > 0 {
		if len(params) > 0 {
			params += "&"
		}

		dateFrom := time.Unix(int64(searchQuery.PublishDateFrom), 0)

		params += "publishDateFrom=" + dateFrom.Format("02.01.2006")
	}

	if searchQuery.PublishDateTo > 0 {
		if len(params) > 0 {
			params += "&"
		}

		dateTo := time.Unix(int64(searchQuery.PublishDateTo), 0)

		params += "publishDateTo=" + dateTo.Format("02.01.2006")
	}

	if len(params) > 0 {
		params += "&"
	}

	params += "searchString=" + url.QueryEscape(searchQuery.SearchString)

	return params
}

func (m GetInfo) Search(searchQuery SearchQuery) string {
	url := "http://www.zakupki.gov.ru/epz/order/quicksearch/search.html?" + m.SearchQueryToParams(searchQuery)
	page, err := http.Get(url)
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
	numBlock := 0
	doc.Find("div").Each(func(i int, s *goquery.Selection) {
		class, _ := s.Attr("class")

		switch class {
		case "registerBox registerBoxBank margBtm20":
			fmt.Println(numBlock)
			OneBlock, _ = s.Html()
			s.Find("td").Each(func(i int, l *goquery.Selection) {
				classx, _ := l.Attr("class")

				switch classx {
				case "tenderTd":

					//fmt.Println(l.Html())

				default:

				}
			})
			massivData = massivData + OneBlock

			fmt.Println(OneBlock)
			numBlock = numBlock + 1
		default:

		}
	})

	//fmt.Println(content, doc.Find(".allRecords").Text())
	returned := massivData + string(header)
	return returned
}
