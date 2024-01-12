package app

import (
	"github.com/Nilfgard13/GOSTORE/app/controller"
)

func (server *Server) InitializeRoutes() {
	server.Router.HandleFunc("/", controller.Home).Methods("GET")
}
