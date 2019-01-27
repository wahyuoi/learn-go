package main

import "net/http"

const USERNAME = "batman"
const PASSWORD = "secret"

func Auth(writer http.ResponseWriter, request *http.Request) bool {
	username, password, ok := request.BasicAuth()
	if !ok {
		writer.Write([]byte("something went wrong"))
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)
	if (!isValid) {
		writer.Write([]byte("wrong username/password"))
		return false
	}

	return true
}

func AllowOnlyGET(writer http.ResponseWriter, request *http.Request) bool {
	if request.Method != "GET" {
		writer.Write([]byte("Only GET is allowed"))
		return false
	}

	return true
}
