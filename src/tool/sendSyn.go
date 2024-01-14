package tool

import (
  "net"
  "math/rand"

  "syn/src/config/model"
  "github.com/google/gopacket"
  "github.com/google/gopacket/layers"
  "golang.org/x/net/ipv4"
)

func SendSYN(c net.PacketConn, flags model.Config) error {
  pTCP := layers.IPProtocolTCP

  ipRawConn, err := ipv4.NewRawConn(c)
  if err != nil {
    return err
  }

  tcp := layers.TCP {
    SrcPort: layers.TCPPort(rand.Intn(65535-1024) + 1024),
    DstPort: layers.TCPPort(flags.Port),
    SYN: true,
    ACK: false,
    Window: 65535,
    Seq: rand.Uint32(),
  }

  opt := gopacket.SerializeOptions {
    FixLengths: true,
  }

  buffer := gopacket.NewSerializeBuffer()
  err = gopacket.SerializeLayers(buffer, opt,
    &tcp,
    gopacket.Payload([]byte{'a', 'b', 's', 'e', 'c', '\n'}),
  )
  if err != nil {
    return err
  }

  b := buffer.Bytes()

  h := &ipv4.Header {
    Version: ipv4.Version,
    Len: ipv4.HeaderLen,
    TotalLen: ipv4.HeaderLen + len(b),
    Protocol: int(pTCP),
    Dst: net.ParseIP(flags.Host).To4(),
    TTL: 64,
  }

  if err := ipRawConn.WriteTo(h, b, nil); err != nil {
    return err
  }

  return nil
}

