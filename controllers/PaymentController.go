package controllers

import (
	// "github.com/ichtrojan/thoth"
	// "html/template"
	// "log"
	// "github.com/rs/xid"
	"io/ioutil"
	"fmt"
	"net/http"
	// "time"
)
func Pay(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
		}

		fmt.Printf("No error, body: %s\n", body)
	}
}