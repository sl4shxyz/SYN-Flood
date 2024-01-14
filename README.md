# This is a tool for sending SYN packets to a specific host

Requirements: ```go version >= 1.21.5 and have administrator privileges (root)```

## Options: 
 - `--help`   Displays a help message
 - `-h`       Host for sending packages
 - `-p`       Port where packets will be sent
 - `-q`       Number of packets sent per second (default = 20)

## Example:
```go run main.go -h host.com -p 80 -q 10```

**Didactic code** 📚
