package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		go requestHandler(conn)

	}
}

func requestHandler(conn net.Conn) {

	defer conn.Close()
	statusOK := "HTTP/1.1 200 OK\r\n\r\n"
	status404 := "HTTP/1.1 404 Not Found\r\n\r\n"
	scanner := bufio.NewScanner(conn)

	if !scanner.Scan() {
		fmt.Printf("[Error while reading data: %s]\n", conn.RemoteAddr())
		os.Exit(1)
	}

	request := scanner.Text()
	splittedReq := strings.Split(request, " ")

	switch {
	case splittedReq[1] == "/":
		conn.Write([]byte(statusOK))
	default:
		conn.Write([]byte(status404))
	}

}
