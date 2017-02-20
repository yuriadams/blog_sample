package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", NicePostHandler)

	log.Fatal(http.ListenAndServe(":8888", router))
}

func NicePostHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  fmt.Fprintf(w, "This is a nice post, don't you think?")
}
