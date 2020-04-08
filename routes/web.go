package routes

import (
	"html/template"
	"log"
	"net/http"
	"github.com/ichtrojan/thoth"
	"twitter-hangouts/controllers"
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	route := mux.NewRouter()
	
	route.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./public"))))
	route.NotFoundHandler = http.HandlerFunc(notFound)
	route.HandleFunc("/api/initiate", controllers.Detail).Methods("POST")
	route.HandleFunc("/api/verify", controllers.Verify).Methods("POST")

	return route
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	logger, _ := thoth.Init("log")

	view, err := template.ParseFiles("views/errors/404.html")

	if err != nil {
		logger.Log(err)
		log.Fatal(err)
	}

	_ = view.Execute(w, nil)
}
