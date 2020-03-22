package utils

import (
	"encoding/json"
	"net/http"
)

type Error struct {
    Mesage string `json:"message"`
}

func Message(status bool, message string) (map[string]interface{}) {
	return map[string]interface{} {"status" : status, "message" : message}
}

func Errors(status bool, message string) (map[string]interface{}) {
	mess := &Error{Mesage: message}
	bs_mess, _ := json.Marshal(mess)
	return Message(status, string(bs_mess))
}

func Respond(w http.ResponseWriter, data map[string] interface{})  {
	// w.Header().Add("Content-Type", "application/json")
    // w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(data)
}