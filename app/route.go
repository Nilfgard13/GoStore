package app

import (
	"net/http"

	"github.com/Nilfgard13/GOSTORE/app/controller"
	"github.com/gorilla/mux"
)

func (server *Server) InitializeRoutes() {

	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controller.Home).Methods("GET")

	staticFileDirectory := http.Dir("./asset/")
	staticFileHandler := http.StripPrefix("/public/", http.FileServer(staticFileDirectory))
	server.Router.PathPrefix("/public/").Handler(staticFileHandler).Methods("GET")
}
