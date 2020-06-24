package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

var (
	logger = log.New(os.Stdout, "127.0.0.1", log.LstdFlags)

	host = flag.String("host", "127.0.0.1", "host to listen on")
	port = flag.Int("port", 2323, "port to listen on")

	blockSize = flag.Int("size", 1024, "block size to read packets on")
)

func main() {
	flag.Parse()

	ip := net.ParseIP(*host)
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: ip, Port: *port})
	if err != nil {
		fmt.Println(err)
		return
	}

	logger.Printf("listening on addr=%s with block size=%d", listener.LocalAddr(), *blockSize)

	data := make([]byte, *blockSize)
	for {
		n, remoteAddr, err := listener.ReadFrom(data)
		if err != nil {
			logger.Fatalf("error during read: %s", err)
		}

		logger.Printf("<%s> %s", remoteAddr, data[:n])
	}
}
