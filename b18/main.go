package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	ActionStudent := func(w http.ResponseWriter, r *http.Request) {
		if !Auth(w, r) {
			return
		}

		if (!AllowOnlyGET(w, r)) {
			return
		}

		if id := r.URL.Query().Get("id"); id != "" {
			OutputJSON(w, SelectStudent(id))
			return
		}
	}

	http.HandleFunc("/student", ActionStudent)


	//server := new(http.Server)
	//server.Addr = "9000"

	fmt.Println("server started at localhost:9000")
	//server.ListenAndServe()
	http.ListenAndServe(":9000", nil)
}

func OutputJSON(writer http.ResponseWriter, student interface{}) {
	res, err := json.Marshal(student)
	if err != nil {
		writer.Write([]byte(err.Error()))
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(res)
	writer.Write([]byte("\n"))
}
