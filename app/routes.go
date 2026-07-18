package app

import "github.com/fthdns/gomart/app/controllers"

func (server *Server) InitializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
