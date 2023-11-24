package routes

import (
	"net/http"
	"ttd/controller"
	"github.com/gorilla/mux"
)

func Routes() *mux.Router{
	// Implement your routes here
	route := mux.NewRouter()
  fs := http.FileServer(http.Dir("./static"))	
	route.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	
	route.HandleFunc("/", controller.Home).Methods("GET")
	route.HandleFunc("/about", controller.About).Methods("GET")
	route.HandleFunc("/upload", controller.Form).Methods("GET")
	route.HandleFunc("/download/{image}", controller.Download).Methods("GET")
	route.HandleFunc("/upload", controller.Upload).Methods("POST")
	return route
}
