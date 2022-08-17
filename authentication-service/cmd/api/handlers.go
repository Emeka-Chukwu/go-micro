package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func (app *Config) Authentication(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	log.Println(r.Body)

	err := app.readJson(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil || user == nil {
		app.errorJSON(w, err, http.StatusBadRequest)

		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)

		return
	}

	payLoad := jsonResponse{Error: false, Message: fmt.Sprintf("Logged in user %s", user.Email), Data: user}

	app.writeJSON(w, http.StatusAccepted, payLoad)
}
