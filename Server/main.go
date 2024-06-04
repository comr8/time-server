package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

const addr = "0.0.0.0:12345" // address and port the server will listen on
const proto = "tcp4"         // TCPv4 protocol

func main() {
	// create a listening socket
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	// close the listening socket when done
	defer listener.Close()

	for {
		// accept incoming connection
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("got new connection: Addr:%v\n", conn.RemoteAddr())
		// handle the connection in a separate goroutine
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	// close the connection when function ends
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		// read data until newline character
		b, err := reader.ReadBytes('\n')
		if err != nil {
			log.Println(err)
			return
		}
		// trim newline characters from the message
		msg := strings.Trim(string(b), "\r\n")
		fmt.Printf("got new msg: %v\n", msg)

		// handle request for sending current time
		if msg == "time" {
			curTime := []byte(time.Now().String() + "\n")
			_, err := conn.Write(curTime)
			if err != nil {
				log.Fatalf("error sending time %v\n", err)
			}
			fmt.Printf("send time: %v\n", string(curTime))
		}
		// handle exit command
		if msg == "exit" {
			err := conn.Close()
			if err != nil {
				log.Fatalf("error closing connection %v\n", err)
			}
			fmt.Printf("Connection for %v closed\n", conn.RemoteAddr())
			return
		}
	}

}
