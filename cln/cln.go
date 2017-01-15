package main

import (
	"log"
	"net"

	"fmt"
	"os"

	"github.com/armon/go-socks5"
	"github.com/inconshreveable/muxado"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid arguments.\n cln.exe <address:port> \nexample usage: cln.exe 127.0.0.1:8484")
		os.Exit(0)
	}

	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	conn, err := net.Dial("tcp", os.Args[1])
	if err != nil {
		log.Println("Cannot connect to target: ", err)
		os.Exit(0)
	}

	sess := muxado.Client(conn, nil)
	for {
		sconn, err := sess.Accept()
		if err != nil {
			log.Println("Can't accept, connection is dead", err)
			break
		}
		go server.ServeConn(sconn)
	}
	// Simple way to keep program running until CTRL-C is pressed.
	//<-make(chan struct{})
}
