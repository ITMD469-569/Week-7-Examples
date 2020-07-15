package main

import (
	"fmt"

	"log"
	"net"
)

func main() {
	// start a TCP server listening on localhost:8088
	connection, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	// run forever, keep listening for connections
	for {
		// accept an incoming connection and create a handle to the connection (conn)
		// you will need to save your conns in the final project, so that you can write to connected clients
		conn, err := connection.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// concurrently handle the connection
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	fmt.Println(c)
	c.Close()
}
