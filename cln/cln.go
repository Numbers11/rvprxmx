package main

import (
	"log"
	"net"

	"github.com/armon/go-socks5"
	"github.com/inconshreveable/muxado"
)

func main() {
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}
	//log.Println("sssssssssss")
	conn, _ := net.Dial("tcp", "127.0.0.1:8484")

	sess := muxado.Client(conn, nil)
	for {
		sconn, err := sess.Accept()
		if err != nil {
			log.Println("Can't accept, connection might be dead", err)
			break
		}
		go server.ServeConn(sconn)
	}
	// Simple way to keep program running until CTRL-C is pressed.
	//<-make(chan struct{})
}
