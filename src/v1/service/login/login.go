package login

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func LoginHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	fmt.Fprint(w, "Welcome2!\n")
}

func IndexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
