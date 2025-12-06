package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

const (
	serverPort = 10000
	serverHost = "149.248.222.193"
)

// echo foo | nc -q 0 149.248.222.193 10000
func main() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverHost, serverPort))
	if err != nil {
		log.Fatal("Connection error", err)
	}
	defer conn.Close()

	data := []byte("Hello")

	_, err = conn.Write(data)
	if err != nil {
		log.Fatal("Write error", err)
	}
	fmt.Printf("Data sent: %s\n", data)

	if tcpConn, ok := conn.(*net.TCPConn); ok {
		tcpConn.CloseWrite()
	}

	echo := make([]byte, len(data))
	nRead, err := io.ReadFull(conn, echo)
	if err != nil {
		fmt.Printf("Read error: %v\n", err)
	}
	fmt.Printf("Data read: %d, %s\n", nRead, echo)
}
