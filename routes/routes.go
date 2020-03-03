package routes

import "github.com/julienschmidt/httprouter"

var router *httprouter.Router

// Init routes handlers
func Init() *httprouter.Router {
	router = httprouter.New()

	initUsers()

	return router
}
