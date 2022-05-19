package main

import (
	"log"
	"net/http"
	"text/template"
)



func f1(w http.ResponseWriter, r *http.Request) {
	// Define the function
	f := func (s string) (string, error) {
		return s + " world!", nil
	}

	t := template.New("f.tmpl")
	// Register the function to the FuncMap
	t.Funcs(template.FuncMap{
		"f": f,
	})
	_, err := t.ParseFiles("f.tmpl")
	if err != nil {
		log.Fatalf("Parse file failed: %v", err)
	}

	t.Execute(w, "hello")
}

func demo(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("t.tmpl", "ul.tmpl")
	if err != nil {
		log.Fatalf("parse tmplate failed: %v", err)
	}

	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmplDemo", demo)

	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatalf("ListenAndServe failed: %v", err)
	}

}