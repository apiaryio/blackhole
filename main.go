package main

import (
	"log"
	"net"
	"os"
)

const (
	MESSAGE = "HTTP/1.1 429 Too Many Requests\r\nConnection: close\r\nContent-Length: 227\r\nContent-Type: text/plain\r\n\r\nWe have detected exceedingly high usage of your mock/proxy server, you are rate limited for the time being. Please reach us at support@apiary.io if you need any help restoring access and please help us understand your use case."
)

func SetupServer() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("Error listening:", err.Error())
	}
	defer l.Close()
	log.Println("Listening on port:" + os.Getenv("PORT"))
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("Error accepting: ", err.Error())
		}
		conn.Write([]byte(MESSAGE))
		conn.Close()
	}
}

func main() {
	SetupServer()
}
