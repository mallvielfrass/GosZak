package main

import (
	"fmt"
	"log"
	"time"
	"strings"
	"strconv"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type GetInfo struct {}

type SearchResult struct {
	Items     []SearchItem
	Total     int64
	Page      int64
	TotalPage int64
}

type SearchItem struct {
	Name         string
	Link         string
	Ids          []string
	Type         string
	Status       string
	Law          string
	Price        string
	Currency     string
	Customer     string
	CustomerLink string
	Description  string
	Lots         []SearchItemLot
	PublishDate  int64
	UpdateDate   int64
	Actions      []SearchItemAction
}

type SearchItemLot struct {
	Name        string
	Description string
	Price       string
	Currency    string
}

type SearchItemAction struct {
	Name string
	Link string
}

func (_ GetInfo) SearchQueryToParams(searchQuery SearchQuery) string {
	params := ""

	for _, item := range searchQuery.LawNumber {
		if len(params) > 0 {
			params += "&"
		}

		switch item {
		case "44-fz":
			params += "fz44=on"
		case "223-fz":
			params += "fz223=on"
		case "pp_rf_615":
			params += "ppRf615=on"
		case "94-fz":
			params += "fz94=on"
		}
	}

	for _, item := range searchQuery.ProcedureStatus {
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

	if len(searchQuery.SortDirection) > 0 {
		if len(params) > 0 {
			params += "&"
		}

		switch searchQuery.SortDirection {
		case "up":
			params += "sortDirection=true"
		case "down":
			params += "sortDirection=false"
		}
	}

	if len(searchQuery.SortBy) > 0 {
		if len(params) > 0 {
			params += "&"
		}

		switch searchQuery.SortBy {
		case "updateDate":
			params += "sortBy=UPDATE_DATE"
		case "publishDate":
			params += "sortBy=PUBLISH_DATE"
		case "price":
			params += "sortBy=PRICE"
		case "relevance":
			params += "sortBy=RELEVANCE"
		}
	}

	if len(searchQuery.CityName) > 0 {
		if len(params) > 0 {
			params += "&"
		}

		switch searchQuery.CityName {
		case "st_petersburg":
			params += "regions=5277347"
		case "moscow":
			params += "regions=5277335"
		}
	}

	if searchQuery.PublishDateFrom > 0 {
		if len(params) > 0 {
			params += "&"
		}

		dateFrom := time.Unix(searchQuery.PublishDateFrom, 0)
		params += "publishDateFrom=" + url.QueryEscape(dateFrom.Format("02.01.2006"))
	}

	if searchQuery.PublishDateTo > 0 {
		if len(params) > 0 {
			params += "&"
		}

		dateTo := time.Unix(searchQuery.PublishDateTo, 0)
		params += "publishDateTo=" + url.QueryEscape(dateTo.Format("02.01.2006"))
	}

	if searchQuery.PageNumber > 0 {
		if len(params) > 0 {
			params += "&"
		}

		params += "pageNumber=" + strconv.FormatInt(searchQuery.PageNumber, 10)
	}

	if len(params) > 0 {
		params += "&"
	}

	params += "searchString=" + url.QueryEscape(searchQuery.SearchString)

	return params
}

func (m GetInfo) Search(searchQuery SearchQuery) (SearchResult, error) {
	var result SearchResult
	var err error

	searchUrl := "http://zakupki.gov.ru/epz/order/quicksearch/search.html?" + m.SearchQueryToParams(searchQuery)

	page, err := http.Get(searchUrl)
	if err != nil {
		return SearchResult{}, err
	}

	defer page.Body.Close()

	if page.StatusCode != 200 {
		return SearchResult{}, fmt.Errorf("Search page status error: %s", page.Status)
	}

	doc, err := goquery.NewDocumentFromReader(page.Body)
	if err != nil {
		return SearchResult{}, err
	}

	currentPageElement := doc.Find(".paginator .page__link_active").First()

	if currentPageElement.Length() > 0 {
		currentPageString := strings.TrimSpace(currentPageElement.Text())
		currentPage, err := strconv.ParseInt(currentPageString, 10, 64)
		if err != nil {
			return SearchResult{}, fmt.Errorf("Error on getting current page number: %s\n", err)
		}

		result.Page = currentPage
	} else {
		result.Page = 1
	}

	totalPageElement := doc.Find(".paginator .page__link").Last()

	if totalPageElement.Length() > 0 {
		totalPageString := strings.TrimSpace(totalPageElement.Text())
		totalPage, err := strconv.ParseInt(totalPageString, 10, 64)
		if err != nil {
			return SearchResult{}, fmt.Errorf("Error on getting total page number: %s\n", err)
		}

		result.TotalPage = totalPage
	} else {
		result.TotalPage = 1
	}

	totalNode := doc.Find(".allRecords > strong").First()
	totalNumber := int64(0)

	if totalNode.Length() > 0 {
		totalComment := totalNode.Get(0).NextSibling

		if len(strings.TrimSpace(totalComment.Data)) > 0 {
			totalString := strings.Split(totalComment.Data, ": ")[1]
			totalNumber, err = strconv.ParseInt(strings.ReplaceAll(totalString, "\u00a0", ""), 10, 64) // there is a no-break space
			if err != nil {
				return SearchResult{}, fmt.Errorf("Error on getting exact total records number: %s\n", err.Error())
			}
		} else {
			totalString := strings.TrimSpace(totalNode.Text())
			totalNumber, err = strconv.ParseInt(strings.ReplaceAll(totalString, "&nbsp;", ""), 10, 64) // there is a no-break space
			if err != nil {
				return SearchResult{}, fmt.Errorf("Error on getting approximate total records number: %s\n", err.Error())
			}
		}
	}

	result.Total = totalNumber
	result.Items = make([]SearchItem, 0)

	doc.Find("div.registerBox.registerBoxBank.margBtm20").Each(func(i int, s *goquery.Selection) {
		var itemStruct SearchItem

		itemTable := s.ChildrenFiltered("table").First()
		itemHeader := itemTable.Find(".descriptTenderTd > dl > dt").First()

		itemHeaderLink := itemHeader.Find("a")
		itemHeaderLinkHref, _ := itemHeaderLink.Attr("href")

		itemUrl, err := url.Parse(itemHeaderLinkHref)
		if err != nil {
			log.Print(err)
			return
		}

		itemUrl.Scheme = "http"
		itemUrl.Host = "zakupki.gov.ru"

		itemName := strings.TrimSpace(itemHeaderLink.Text())
		itemHeader.Remove()

		itemOrganization := itemTable.Find(".descriptTenderTd > dl > .nameOrganization").First()

		organizationLink := itemOrganization.Find("a").First()
		organizationLinkHref, _ := organizationLink.Attr("href")

		organizationUrl, err := url.Parse(organizationLinkHref)
		if err != nil {
			log.Print(err)
			return
		}

		//fix relative urls
		organizationUrl.Scheme = "http"
		organizationUrl.Host = "zakupki.gov.ru"

		organizationName := strings.TrimSpace(organizationLink.Text())
		itemOrganization.Remove()

		itemIdNode := itemTable.Find(".descriptTenderTd > dl > dd.padTop10 > dl.greyText.margTop0.padTop8").First()
		itemIdNode.Find("script").Remove()
		itemId := strings.ReplaceAll(itemIdNode.Text(), " ", "")
		itemIdNode.Parent().Remove()

		itemIds := strings.Split(itemId, "\n")
		tempItemIds := make([]string, 0)

		for i := 0; i < len(itemIds); i++ {
			itemIds[i] = strings.TrimSpace(itemIds[i])

			if len(itemIds[i]) > 0 {
				tempItemIds = append(tempItemIds, itemIds[i])
			}
		}

		itemIds = tempItemIds

		itemDescription := strings.TrimSpace(itemTable.Find(".descriptTenderTd > dl > *").First().Text())
		itemType := strings.TrimSpace(itemTable.Find(".tenderTd > dl > dt > strong").First().Text())
		itemStatusSlice := strings.Split(itemTable.Find(".tenderTd > dl > dt > span.noWrap").First().Text(), "/")
		itemStatus := strings.TrimSpace(itemStatusSlice[0])
		itemLaw := strings.TrimSpace(itemStatusSlice[1])
		itemPriceSlice := strings.Split(itemTable.Find(".tenderTd > dl > dd .fractionalNumber").First().Parent().Text(), ",")

		for i2, priceItem := range itemPriceSlice {
			itemPriceSlice[i2] = strings.TrimSpace(priceItem)
		}

		itemPrice := strings.Join(itemPriceSlice, ",")
		itemCurrency := strings.TrimSpace(itemTable.Find(".tenderTd > dl > dd > .currency").First().Text())

		itemStruct.Lots = make([]SearchItemLot, 0)
	
		itemTable.Find(".lotsInfo .descriptTenderTd").Each(func(i2 int, s2 *goquery.Selection) {
			var lotInfo SearchItemLot

			lotDesciptionElem := s2.Find("dl > dt").First()
			lotNameElem := lotDesciptionElem.ChildrenFiltered("strong").First()
			lotPriceElem := s2.Find("dl > dt > i > strong").First()

			lotInfo.Name = strings.TrimSpace(lotNameElem.Text())
			lotNameElem.Remove()

			lotInfo.Description = strings.TrimSpace(lotDesciptionElem.Text())
			lotInfo.Price = strings.TrimSpace(lotPriceElem.Text())
			lotInfo.Currency = strings.TrimSpace(lotPriceElem.Get(0).NextSibling.Data)

			itemStruct.Lots = append(itemStruct.Lots, lotInfo)
		})

		amountNodes := itemTable.Find(".amountTenderTd > ul > li > label")
		publishDateString := strings.TrimSpace(amountNodes.Get(0).NextSibling.Data)
		updateDateString := strings.TrimSpace(amountNodes.Get(1).NextSibling.Data)

		publishDate, err := time.Parse("02.01.2006", publishDateString)
		if err != nil {
			log.Print(err)
			return
		}

		updateDate, err := time.Parse("02.01.2006", updateDateString)
		if err != nil {
			log.Print(err)
			return
		}

		reportBox := itemTable.Next()
		reportBoxList := reportBox.Find("ul > ul").First()

		reportBoxList.Find("a").Each(func(i2 int, s2 *goquery.Selection) {
			reportHref, reportHrefExists := s2.Attr("href")
			if !reportHrefExists {
				s2.Remove()
				return
			}

			reportUrl, err := url.Parse(reportHref)
			if err != nil {
				log.Print(err)
				return
			}

			//fix relative urls
			reportUrl.Scheme = "http"
			reportUrl.Host = "zakupki.gov.ru"

			var reportItem SearchItemAction
			reportItem.Name = strings.TrimSpace(s2.Text())
			reportItem.Link = reportUrl.String()

			itemStruct.Actions = append(itemStruct.Actions, reportItem)
		})

		itemStruct.Name = itemName
		itemStruct.Link = itemUrl.String()
		itemStruct.Ids = itemIds
		itemStruct.Type = itemType
		itemStruct.Status = itemStatus
		itemStruct.Law = itemLaw
		itemStruct.Price = itemPrice
		itemStruct.Currency = itemCurrency
		itemStruct.Customer = organizationName
		itemStruct.CustomerLink = organizationUrl.String()
		itemStruct.Description = itemDescription
		itemStruct.PublishDate = publishDate.Unix()
		itemStruct.UpdateDate = updateDate.Unix()

		result.Items = append(result.Items, itemStruct)
	})

	return result, nil
}
