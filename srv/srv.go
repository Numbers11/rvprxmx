package main

import (
	"fmt"
	"log"
	"net"
	"crypto/tls"
	"time"

	"github.com/inconshreveable/muxado"
)

var cfg *config

var clientjson []byte

var clients = make(map[muxado.Session]*client)

const banner = `
 _ ____   ___ __  _ ____  ___ __ ___ __  __
| '__\ \ / / '_ \| '__\ \/ / '_ ` + "`" + ` _ \\ \/ /
| |   \ V /| |_) | |   >  <| | | | | |>  < 
|_|    \_/ | .__/|_|  /_/\_\_| |_| |_/_/\_\
           |_|                 !TRUMP!                     
`

//Listen create a listener and serve on it
func listen() error {
	cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
	        return err
	}

	tlsconfig := &tls.Config{Certificates: []tls.Certificate{cer}}

	l, err := tls.Listen("tcp", ":"+cfg.CnCPort, tlsconfig)

	//l, err := net.Listen("tcp", ":"+cfg.CnCPort)

	if err != nil {
		return err
	}

	log.Println("Started CnC server at port " + cfg.CnCPort)

	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("TCP accept failed: %s\n", err)
			continue
		}
		go handle(conn)
	}
}

//handle a new client connection
func handle(conn net.Conn) {
	// Setup server side of muxado
	session := muxado.Server(conn, nil)
	defer session.Close()

	//Setup tcp server on random free port
	listener, err := net.Listen("tcp", ":")
	if err != nil {
		log.Println("Error starting Socks listener", err)
		return
	}

	//create new Client
	client := newClient(session, listener, cfg.SocksUsername, randomString(6))
	clients[session] = client
	go client.listen()
	log.Printf("Started new SOCKS listener at port %v, auth %v:%v\n", listener.Addr().String(), client.Username, client.Password)

	//blocks until the muxado connnection is closed
	client.wait()

	//delete client
	delete(clients, session)
}

func main() {
	fmt.Println(banner)
	var err error
	cfg, err = loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	//should/could possibly be lower or smth
	schedule(updatejson, 1000*time.Millisecond)
	//schedule(func() { log.Println(clientjson) }, 2000*time.Millisecond)

	go startHTTP(cfg.HTTPPort)
	log.Println("Started HTTP server at port " + cfg.HTTPPort)

	//run main tcp server
	listen()
}
