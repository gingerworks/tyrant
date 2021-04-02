package core

import (
	"fmt"
	"net/http"
)

type RouteURL struct {

}

func RegisterRoutes(mux *http.ServeMux)  {
  mux.HandleFunc("/", indexHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hi there, you are looking for %s, yes?", r.URL.Path[1:])
}
