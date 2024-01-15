package app

import (
	"github.com/Nilfgard13/GOSTORE/app/controller"
	"github.com/gorilla/mux"
)

func (server *Server) InitializeRoutes() {

	server.Router = mux.NewRouter()

	server.Router.HandleFunc("/", controller.Home).Methods("GET")
}
