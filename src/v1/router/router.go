package router

import (
	"github.com/julienschmidt/httprouter"
	login "github.com/metabbe3/Knoxsbite-Backend/src/v1/service/login"
)

func InitRouter(r *httprouter.Router) {
	setLoginRoutes(r)
}

func setLoginRoutes(r *httprouter.Router) {
	r.GET("/", login.IndexHandler)
	r.POST("/login", login.LoginHandler)
}
