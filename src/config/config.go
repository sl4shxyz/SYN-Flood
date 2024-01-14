package config

import (
  "flag"
  "errors"
  "syn/src/config/model"
)

func ParseFlags() (*model.Config, error) {
  hostFlag := flag.String("h", "", "Target for sending SYN packets")
  portFlag := flag.Int("p", 0, "Port for sending packages")
  quantityFlag := flag.Int("q", 20, "Number of packets sent per second")

  flag.Parse()

  if *hostFlag == "" || *portFlag == 0 {
    return nil, errors.New("Host and port are required")
    flag.PrintDefaults()
  }

  return &model.Config {
    Host: *hostFlag,
    Port: *portFlag,
    Quantity: *quantityFlag,
  }, nil

}

func PrintDefaults() {
  flag.PrintDefaults()
}
