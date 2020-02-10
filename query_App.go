package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("http://www.qq.com")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	fmt.Println(res.Body)
	bytes, _ := ioutil.ReadAll(res.Body)

	out, _ := iconv.ConvertString(string(bytes), "gbK", "utf-8")
	fmt.Println(out)
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(out))
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("div").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band, _ := s.Find("a").Html()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})
}

func main() {
	ExampleScrape()
}
