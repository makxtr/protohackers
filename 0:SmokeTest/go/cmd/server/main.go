package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const (
	defaultPort = 2000
)

func RunServer(l net.Listener) {
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("New connection from %s", conn.RemoteAddr())
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Printf("Handling connection from %s", conn.RemoteAddr())

	n, err := io.Copy(conn, conn)
	log.Printf("Copied %d bytes, error: %v", n, err)

	if err != nil && err != io.EOF {
		log.Printf("Error copy data: %v", err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = fmt.Sprintf("%d", defaultPort)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	log.Printf("Server listening on %s", listener.Addr().String())
	RunServer(listener)
}
