package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type NewsAggPage struct {
	Title string
	News  string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Export web design by TaHa")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Amazing news aggregator", News: "some news"}
	t, _ := template.ParseFiles("basictemplating.html")
	t.Execute(w, p)

}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about/", aboutHandler)
	http.HandleFunc("/agg/", newsAggHandler)

	_ = http.ListenAndServe(":8000", nil)

}
