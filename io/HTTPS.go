package io

import (
	"net/http"
	"os"
)

var user = os.Getenv("user")
var userID = "/users/" + user + "/"
var dashboard = userID + "dashboard/"

func featureHandler() {
	http.HandleFunc("/features/", func(w http.ResponseWriter, r *http.Request) {

	})
}

func moduleHandler() {
	http.HandleFunc(userID, func(w http.ResponseWriter, r *http.Request) {

	})
}
func updateHandlerModule() {
	http.HandleFunc(dashboard, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
		case http.MethodDelete:
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)

		}
	})
}
