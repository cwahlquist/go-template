package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerTemplate(w http.ResponseWriter, r *http.Request) {
	title := "Circadence GoLang Template"

	from := ""
	if r.URL != nil {
		from = r.URL.String()
	}
	if from != "/favicon.ico" {
		log.Printf("title: %s\n", title)
	}

	fmt.Fprintf(w, "Hello Bob, from:  "+title+"\n")
}

func main() {
	http.HandleFunc("/", handlerTemplate)
	http.ListenAndServe(":8080", nil)
}
