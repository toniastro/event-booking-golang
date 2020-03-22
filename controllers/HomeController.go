package controllers

import (
	// "html/template"
	"net/http"
	"time"
	"twitter-hangouts/views"
	"twitter-hangouts/models"
	"encoding/json"
	u "twitter-hangouts/utils"

)

var results []string

var view *views.View


func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

var Detail = func(w http.ResponseWriter, r *http.Request) {

	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Headers","Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	detail := &models.Details{}
	err := json.NewDecoder(r.Body).Decode(detail) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := detail.Create() //Create account
	u.Respond(w, resp)
}
var Verify = func(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	detail := &models.Payload{}
	err := json.NewDecoder(r.Body).Decode(detail) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := detail.Confirm() //Create account
	u.Respond(w, resp)
}

