package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args)-1 >= 1 {
		fmt.Printf("HELP:\n0. The client automatically connects to the server (if available) via ip: localhost:12345\n1. Enter any text to send to the server. By default, the server processes only the time command\n2. If there is no response from the server within 3 seconds, you will be prompted to enter a new text\n3. Enter exit to disconnect from the server\n")
		os.Exit(1)
	}
	// Connect to a network service
	conn, err := net.Dial("tcp4", "localhost:12345")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Printf("Succesfully connected to: %v\b with ip: %v\n", conn.RemoteAddr(), conn.LocalAddr())
	// Set the timeout for reading the response from the server to 3 seconds
	conn.SetReadDeadline(time.Now().Add(3 * time.Second))

	readerResponse := bufio.NewReader(conn)
	for {
		// Reading the line to be sent to the server
		fmt.Print("Enter text to send: ")
		text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if text == "exit\n" {
			fmt.Println("Bye!")
			return
		}

		// Sending text to the server
		_, err := conn.Write([]byte(text))
		if err != nil {
			log.Fatal(err)
		}

		// Reading the response from the server
		response, err := readerResponse.ReadString('\n')
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				fmt.Println("Timed out waiting for a response from the server")
			} else {
				log.Fatal(err)
			}
		}
		if len(response) != 0 {
			fmt.Println("Response from the server:", response)
		}

	}
}
