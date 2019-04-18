package main

import (
	"fmt"
	"time"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type GetInfo struct {}

type SearchItemBlock struct {
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

		switch item {
		case "94-fz":
			params += "fz94=on"
		case "pp_rf_615":
			params += "ppRf615=on"
		case "223-fz":
			params += "fz223=on"
		case "44-fz":
			params += "fz44=on"
		}
	}

	for _, item := range searchQuery.ProcedureType {
		if len(params) > 0 {
			params += "&"
		}

		switch item {
		case "applicationSubmission":
			params += "af=on"
		case "commissionWork":
			params += "ca=on"
		case "procedureCompleted":
			params += "pc=on"
		case "procedureAborted":
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

func (m GetInfo) Search(searchQuery SearchQuery) (string, error) {
	url := "http://zakupki.gov.ru/epz/order/quicksearch/search.html?" + m.SearchQueryToParams(searchQuery)

	page, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer page.Body.Close()

	if page.StatusCode != 200 {
		return "", fmt.Errorf("status code error: %d %s", page.StatusCode, page.Status)
	}

	doc, err := goquery.NewDocumentFromReader(page.Body)
	if err != nil {
		return "", err
	}

	itemBlocks := ""
	header, _ := goquery.OuterHtml(doc.Find(".allRecords"))

	doc.Find("div.registerBox.registerBoxBank.margBtm20").Each(func(i int, s *goquery.Selection) {
		itemBlock, _ := s.Html()
		itemBlocks += itemBlock

		/* s.Find("td.tenderTd").Each(func(i int, l *goquery.Selection) {
			//fmt.Println(l.Html())
		}) */
	})

	returned := itemBlocks + header
	return returned, nil
}
