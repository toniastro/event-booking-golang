package controllers

import (
	"net/http"
	"twitter-hangouts/views"
	"twitter-hangouts/models"
	"encoding/json"
	u "twitter-hangouts/utils"
)

var results []string

var view *views.View

var Detail = func(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	detail := &models.Details{}
	err := json.NewDecoder(r.Body).Decode(detail) 
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := detail.Create() 
	u.Respond(w, resp)
}
var Verify = func(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	detail := &models.Payload{}
	err := json.NewDecoder(r.Body).Decode(detail) 
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := detail.Confirm() 
	u.Respond(w, resp)
}

