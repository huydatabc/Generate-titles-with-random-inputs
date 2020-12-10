package main

import (
	"GenNameFromKey/route"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	er2 := route.Route(r)
	if er2 != nil {
		panic(er2)
	}

	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":416", r))
}
