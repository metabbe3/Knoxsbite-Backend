package Knoxsbite

import (
	"github.com/julienschmidt/httprouter"
	router "github.com/metabbe3/Knoxsbite-Backend/src/v1/router"
)

func InitKnoxsbite(r *httprouter.Router) {
	router.InitRouter(r)
}
