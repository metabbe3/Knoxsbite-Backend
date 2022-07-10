package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	Knoxsbite "github.com/metabbe3/Knoxsbite-Backend/src/Knoxsbite"
)

func main() {
	r := httprouter.New()
	log.Println("Init Knoxsbite")
	Knoxsbite.InitKnoxsbite(r)
	fmt.Println(http.ListenAndServe(":8080", r))
	//http.Handle("/", r)
}
