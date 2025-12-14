package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"prime/cmd/internal/prime"
	"prime/cmd/internal/request"
)

type Response struct {
	Method string `json:"method,omitempty"`
	Prime  bool   `json:"prime"`
	Error  string `json:"error,omitempty"`
}

func main() {

	listener, err := net.Listen("tcp", ":2002")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Printf("Server listening on %s", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	log.Println("New connection from:", conn.RemoteAddr())

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadBytes('\n')
		if err != nil {
			if err.Error() != "EOF" {
				log.Println("Connection closed or read error:", err)
			}
			return
		}
		log.Printf("Received request: %s", string(message))

		var r Response
		req, err := request.CreateReq(message)
		if err != nil {
			log.Printf("Request error: %v", err)
			r = Response{Error: "Malformed response"}
			sendResponse(conn, r)
			return
		} else {
			isPrime := prime.Prime(int(*req.Number))
			r = Response{Method: req.Method, Prime: isPrime}
			log.Printf("Processing valid request: %+v, Prime: %v", req, isPrime)
			sendResponse(conn, r)
		}
	}
}

func sendResponse(conn net.Conn, r Response) {
	data, err := json.Marshal(r)
	if err != nil {
		log.Printf("Marshal error: %v", err)
		return
	}
	_, err = conn.Write(append(data, '\n'))
	if err != nil {
		log.Printf("Write error: %v", err)
	}
}
