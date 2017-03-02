package main

import (
	"fmt"
	"net"
	"os"
)

const (
	CONN_TYPE = "tcp"
	MESSAGE   = "HTTP/1.1 429 Too Many Requests\r\nConnection: close\r\nContent-Length: 226\r\nContent-Type: text/plain\r\n\r\nWe have detected exceedingly high usage of your mock/proxy server, you are rate limited for the time being. Please reach us at support@apiary.io if you need any help restoring access and please help us understand your use case"
)

func main() {

	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, ":"+os.Getenv("PORT"))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on port:" + os.Getenv("PORT"))
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		conn.Write([]byte(MESSAGE))
		// Close the connection when you're done with it.
		conn.Close()
	}
}
