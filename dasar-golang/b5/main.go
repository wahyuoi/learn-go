package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

type M map[string]interface{}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Batman"}
		
		var tmpl = template.Must(template.ParseFiles(
			path.Join("b5","views","index.html"),
			path.Join("b5","views","_header.html"),
			path.Join("b5","views","_message.html"),
		))
		var err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Batman"}
		var tmpl = template.Must(template.ParseFiles(
			path.Join("b5","views","about.html"),
			path.Join("b5","views","_header.html"),
			path.Join("b5","views","_message.html"),
		))

		var err = tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
