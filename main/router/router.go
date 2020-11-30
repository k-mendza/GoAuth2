package router

import (
	"GoAuth2/main/controllers"
	"github.com/julienschmidt/httprouter"
)

func CreateRoutes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/users/:username", controllers.GetUserByUsername)
	return router
}
