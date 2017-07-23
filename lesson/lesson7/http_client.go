package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func main() {
	url := "http://daily.zhihu.com"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}
	//io.Copy(os.Stdout, resp.Body)
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("src")
		if ok {
			fmt.Println(link)
		} else {
			fmt.Println("src not found.")
		}
	})
}
