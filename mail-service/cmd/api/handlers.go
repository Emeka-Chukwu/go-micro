package main

import (
	"log"
	"net/http"
)

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	type mailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var requestPayload mailMessage
	log.Println(r.Body)
	err := app.readJson(w, r, &requestPayload)
	if err != nil {
		log.Println(err, requestPayload)
		log.Printf("checking")
		app.errorJSON(w, err)
		return
	}
	log.Printf("passed")
	log.Println(requestPayload)
	msg := Message{
		From:    requestPayload.From,
		To:      requestPayload.To,
		Subject: requestPayload.Subject,
		Data:    requestPayload.Message,
	}
	err = app.Mailer.SendSMTPMessage(msg)

	if err != nil {
		app.errorJSON(w, err)
		log.Println(err, "failed here")
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Sent to " + requestPayload.To,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
