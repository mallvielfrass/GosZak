package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func get(params string) string {
	url := "http://www.zakupki.gov.ru/epz/order/quicksearch/search.html?" + params
	resp, err := http.Get(url) // выполняем запрос
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body) // считываем результат
	if err != nil {
		log.Fatalln(err)
	}

	return string(b)
}
