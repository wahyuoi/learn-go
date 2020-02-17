package main

import (
	"fmt"
	"net/http"
)

func main() {
	handlerIndex := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	}

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/data", func(writer http.ResponseWriter, r *http.Request) {
		writer.Write([]byte("Hello Data"))
	})

	addr := ":9000"
	fmt.Print("server started at " + addr)

	http.ListenAndServe(addr, nil)
}
