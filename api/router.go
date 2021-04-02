package api

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  libvirt "github.com/libvirt/libvirt-go"
)

func RegisterRoutes(mux *http.ServeMux)  {
  mux.HandleFunc("/api/", indexHandler)
  mux.HandleFunc("/api/domain/", domainHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "My path %s", r.URL.Path)
}

func domainHandler(w http.ResponseWriter, r *http.Request)  {
  for _, dom := range ListAll() {
    name, err1 := dom.GetName()
    if err1 != nil {
        log.Print("Error Listing: %s", err1)
    }
    info, err2 := dom.GetInfo()
    if err2 != nil {
        log.Print("Error Listing: %s", err2)
    }
    fmt.Fprintf(w, "VM: %s", name)
    if info.State == 1 {
      fmt.Fprintf(w, ", State: Running")
    } else {
      fmt.Fprintf(w, ", State: Shutdown")
    }
    fmt.Fprintf(w, ", %dGB RAM", info.Memory / 1024)
    fmt.Fprintf(w, ", %d CPUs", info.NrVirtCpu)
    jdog, err3 := json.Marshal(info)
    if err3 != nil {
        log.Print("Error Listing: %s", err3)
    }
    fmt.Fprintf(w, ", JSON: %s\n", string(jdog))
  }
  fmt.Fprintf(w, "\n\nTesting a thing: %d", libvirt.DOMAIN_NOSTATE)

}
