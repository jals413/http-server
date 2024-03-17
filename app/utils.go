package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

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

	if len(splittedReq) != 3 {
		fmt.Println("Invalid request")
		os.Exit(1)
	}
	if splittedReq[1] == "/" {
		conn.Write([]byte(statusOK))
		return
	}
	if !strings.HasPrefix(splittedReq[1], "/echo/") {
		conn.Write([]byte(status404))
		return
	}
	param := strings.Replace(splittedReq[1], "/echo/", "", 1)
	conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(param), param)))

}
