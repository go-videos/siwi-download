package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"siwi-download/libs"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Example()
	// E2()

	libs.Download("https://cdn.npm.taobao.org/dist/node/v14.16.0/node-v14.16.0.pkg", "/Volumes/ssd/volumes/code/go-videos/siwi-download/storage")
}

func E2() {
	html := "file: https://cdnlib.tv/videos/1234.m3u8"
	reg := regexp.MustCompile(`https(\S)*m3u8`)

	fmt.Printf("%q\n", reg.Find([]byte(html)))
}
func Example() {
	res, err := http.Get("https://www.bootcdn.cn/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	type Pkg struct {
		title string
		desc  string
	}
	doc.Find(".package").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		title := s.Find("h4.package-name").Text()
		desc := s.Find("p.package-description").Text()
		p := Pkg{title: title, desc: desc}
		fmt.Printf("%+v\n", p)
	})

	html := "file: https://cdnlib.tv/videos/1234.m3u8"
	reg := regexp.MustCompile(`https(\S)*m3u8`)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reg.MatchString(html))
	fmt.Printf("%+v\n", reg)
}
