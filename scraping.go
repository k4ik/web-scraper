package main 

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url:= ""

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Erro ao ler o documento HTTP:" , err)
	}

	document.Find("p").Each(func(index int, element *goquery.Selection) {
		fmt.Println(element.Text())
	})
}