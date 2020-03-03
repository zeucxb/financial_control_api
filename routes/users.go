package routes

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func initUsers() {
	basePath := "/users"

	router.POST(basePath, createUser)
}

func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "New User")
}
