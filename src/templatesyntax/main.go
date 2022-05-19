package main

import (
	"log"
	"net/http"
	"text/template"
)



type User struct {
	Name string
	Gender string
	Age int
}

func test(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("test.tmpl")
	if err != nil {
		log.Fatalf("ParseFiles failed: %v", err)
	}

	u1 := User{
		Name: "y",
		Gender: "男",
		Age: 26,
	}

	m1 := map[string]interface{}{
		"name": "y",
		"gender": "男",
		"age": 26,
	}

	s1 := []string{"a", "b"}

	t.Execute(
		w,
		map[string]interface{}{
			"m1": m1,
			"u1": u1,
			"s1": s1,
		},
	)
}

func main() {
	http.HandleFunc("/", test)
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatalf("listenAndServe failed: %v", err)
	}
}
