package main

import (
	"io"
	"log"
	"net"
)

func main() {
	// dial a tcp connection to a server running on localhost:8088
	conn, err := net.Dial("tcp", "localhost:8088")
	if err != nil {
		log.Fatal(err)
	}

	// write text to the server
	io.WriteString(conn, "Hello Server")
	conn.Close()
}
