package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {

	os.Remove("../server/out.txt")

	// Connect to server on localhost:8080
	pi_ip := os.Getenv("pi_ip")
	//conn, err := net.Dial("tcp", "localhost:8082")
	conn, err := net.Dial("tcp", pi_ip+":8082")
	if err != nil {
		log.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Get command-line arguments
	args := os.Args[1:]

	// Convert args to string and send over network
	argsString := strings.Join(args, " ") + "\n"
	// prepend "dbt"
	argsString = "dbt " + argsString
	log.Println("Sending:", argsString)
	_, err = conn.Write([]byte(argsString))
	if err != nil {
		log.Println("Error sending data:", err)
		return
	}

	// wait for response
	for {
		_, err := os.Stat("../server/out.txt")
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}

	// Read response from server
	file, err := os.Open("../server/out.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// create a scanner
	scanner := bufio.NewScanner(file)

	log.Println("Response:")
	// read line by line
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
