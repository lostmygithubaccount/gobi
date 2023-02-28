package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os/exec"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read data sent by client
	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	// Convert data to slice of strings
	args := strings.Fields(data)
	log.Println(args)

	// Execute command and capture output
	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	// Write output to file
	err = ioutil.WriteFile("out.txt", output, 0644)
	if err != nil {
		fmt.Println("Error writing output to file:", err)
		return
	}

	fmt.Println("Command output written to out.txt")
}

func main() {
	// Listen for incoming connections on port 8080
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server running on port 8082...")

	// Accept incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle each connection in a separate goroutine
		go handleConnection(conn)
	}
}
