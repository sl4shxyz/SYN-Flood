package config

import (
  "flag"
  "os"

  "syn/src/config/model"
)

func ParseFlags() (*model.Config) {
  hostFlag := flag.String("h", "", "Target for sending SYN packets")
  portFlag := flag.Int("p", 0, "Port for sending packages")
  quantityFlag := flag.Int("q", 20, "Number of packets sent per second")

  flag.CommandLine.Parse(os.Args[1:])

  return &model.Config {
    Host: *hostFlag,
    Port: *portFlag,
    Quantity: *quantityFlag,
  }
}

