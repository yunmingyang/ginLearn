package main

import (
	"html/template"
	"log"
	"net/http"
	t_template "text/template"
)

func preventXSS(w http.ResponseWriter, r *http.Request) {
	t, err := t_template.ParseFiles("pxss.tmpl")
	if err != nil {
		log.Printf("template parsefiles failed: %v\n", err)
		return
	}

	t.Execute(w, "<script>alert(1)</script>")
}

func XSS(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"xsst": func(s string) template.HTML {
			return template.HTML(s)
		},
	}).ParseFiles("xss.tmpl")
	if err != nil {
		log.Printf("parse file failed: %v\n", err)
		return
	}

	t.Execute(w, map[string]string{
		"str1": "alert(1)",
		"str2": "<a href='https://www.baidu.com'>百度</a>",
	})
}

func main() {
	http.HandleFunc("/pxss", preventXSS)
	http.HandleFunc("/xss", XSS)

	if err := http.ListenAndServe(":9091", nil); err != nil {
		log.Fatalf("ListenAndServe failed: %v\n", err)
	}
}
