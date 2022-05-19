package main

import (
	"log"
	"net/http"
	"text/template"
)



func main() {
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatalf("listen and serve failed: %v\n", err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// parse template
	t, err := template.ParseFiles("hello.tmpl")
	if err != nil {
		log.Fatalf("Parse faied: %v", err)
	}
	// render template
	err = t.Execute(w, "y")
	if err != nil {
		log.Fatalf("Render failed: %v", err)
	}
}