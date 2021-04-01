package main

import (
	"log"
	"net/http"
	"github.com/gingerworks/tyrant/core"
	"github.com/gingerworks/tyrant/api"
)

func serveTyrant()  {
	mux := http.NewServeMux()
	core.RegisterRoutes(mux)
	api.RegisterRoutes(mux)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func main() {
	serveTyrant()
}
