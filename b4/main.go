package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		var filepath = path.Join("b4", "views", "index.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		var data = map[string]interface{}{
			"title": "Learning Golang",
			"name":  "Batman",
		}

		err = tmpl.Execute(writer, data)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir(path.Join("b4", "assets")))))

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
