package api

import (
  "fmt"
  "log"
	"github.com/gingerworks/tyrant/core"
  libvirt "github.com/libvirt/libvirt-go"
)

func ListAll() []libvirt.Domain {
  connstring := fmt.Sprintf("qemu+ssh://%s@%s/system",core.Config["user"] , core.Config["host"])
  conn, err := libvirt.NewConnect(connstring)
  if err != nil {
      log.Print("Error Connecting: %s", err)
  }
  defer conn.Close()

  doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_PERSISTENT)
  if err != nil {
      log.Print("Error Listing: %s", err)
  }
  return doms
}
