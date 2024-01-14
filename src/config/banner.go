package config

import (
  "fmt"
  "syn/src/config/model"
)


func Banner(model *model.Config) string {
  banner := fmt.Sprintf(`
    ======================================================
     Host:              %s
     Port:              %d
     Quantity:          %d
    ======================================================
  `, model.Host, model.Port, model.Quantity)

   return banner

}

func Help() string {
  help := `HELP:
  -h string
    	Target for sending SYN packets
  -p int
    	Port for sending packages
  -q int
    	Number of packets sent per second (default 20)
  `
  return help
}
