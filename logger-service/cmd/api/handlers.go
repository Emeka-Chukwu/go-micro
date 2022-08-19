package main

import "net/http"

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	// read json into var
	// var requestPayload JSONPayload
	// _ = app.readJSON
}
