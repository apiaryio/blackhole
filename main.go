package main

import (
	"fmt"
	"net"
	"os"
)

const (
	MESSAGE = "HTTP/1.1 429 Too Many Requests\r\nConnection: close\r\nContent-Length: 227\r\nContent-Type: text/plain\r\n\r\nWe have detected exceedingly high usage of your mock/proxy server, you are rate limited for the time being. Please reach us at support@apiary.io if you need any help restoring access and please help us understand your use case."
)

func SetupServer() {
	l, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on port:" + os.Getenv("PORT"))
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		conn.Write([]byte(MESSAGE))
		conn.Close()
	}
}

func main() {
	SetupServer()
}
