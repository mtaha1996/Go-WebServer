package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

/*type SiteMapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}*/

// Simplify the structs
type SiteMapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>Keywords"`
	Locations []string `xml:"url>loc"`
}


func main() {
	var s SiteMapIndex
	var n News

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	//stringBody := string(bytes)
	//fmt.Println(stringBody)
	resp.Body.Close()


	xml.Unmarshal(bytes, &s)

	//fmt.Println(s.Locations)

	for _, Location := range s.Locations {
		//fmt.Printf("\n%s", Location)

		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

		resp.Body.Close()

	}



}
