package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type SiteMapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

type NewsAggPage struct {
	Title string
	News  map[string]NewsMap
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Export web design by TaHa")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {

	var s SiteMapIndex
	var n News
	news_map := make(map[string]NewsMap)

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	//stringBody := string(bytes)
	//fmt.Println(stringBody)
	_ = resp.Body.Close()

	_ = xml.Unmarshal(bytes, &s)

	//fmt.Println(s.Locations)

	for _, Location := range s.Locations {
		//fmt.Printf("\n%s", Location)

		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		_ = xml.Unmarshal(bytes, &n)

		for idx, _ := range n.Titles {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
		_ = resp.Body.Close()
	}

	p := NewsAggPage{Title: "Amazing news aggregator", News: news_map}
	t, _ := template.ParseFiles("basictemplating.html")
	t.Execute(w, p)

}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about/", aboutHandler)
	http.HandleFunc("/agg/", newsAggHandler)

	_ = http.ListenAndServe(":8000", nil)

}
