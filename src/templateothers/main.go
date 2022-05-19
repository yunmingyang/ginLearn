package main

import (
	"log"
	"net/http"
	"text/template"
)



func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("index.tmpl").Delims("{[", "]}").ParseFiles("index.tmpl")
	if err != nil {
		log.Fatalf("template creation failed: %v", err)
	}

	t.Execute(w, "aaaa")
}

func main() {
	http.HandleFunc("/index", index)

	if err := http.ListenAndServe(":9091", nil); err != nil {
		log.Fatalf("ListenAndServe failed: %v", err)
	}
}