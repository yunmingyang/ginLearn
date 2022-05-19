package main

import (
	"log"
	"net/http"
	"text/template"
)



func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/base.tmpl", "template/index.tmpl")
	if err != nil {
		log.Fatalf("pares failed: %v", err)
	}

	t.Execute(w, "index")
}

func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/base.tmpl", "template/home.tmpl")
	if err != nil {
		log.Fatalf("pares failed: %v", err)
	}

	t.Execute(w, "home")
}


func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)

	if err := http.ListenAndServe(":9091", nil); err != nil {
		log.Fatalf("ListenAndServe failed: %v", err)
	}
}