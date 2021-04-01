package api

import (
  "fmt"
  "net/http"
)

func RegisterRoutes(mux *http.ServeMux)  {
  mux.HandleFunc("/api/", indexHandler)
  mux.HandleFunc("/api/vm/", vmHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "My path %s", r.URL.Path)
}

func vmHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "VM go %s", r.URL.Path[len("/api/vm/"):])
}
