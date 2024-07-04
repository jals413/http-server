package main

import (
	"fmt"
	"net"
	"strings"
)

func requestHandler(conn net.Conn) {

	defer conn.Close()

	statusOK := "HTTP/1.1 200 OK\r\n\r\n"
	status404 := "HTTP/1.1 404 Not Found\r\n\r\n"

	buffer := make([]byte, 1024)
	conn.Read(buffer)

	splitedHeader := strings.Split(string(buffer), "\r\n")
	splitedRequestLine := strings.Split(splitedHeader[0], " ")
	splitedUserAgent := strings.Split(splitedHeader[2], " ")

	if splitedRequestLine[1] == "/" {
		conn.Write([]byte(statusOK))
		return
	} else if strings.Split(splitedRequestLine[1], "/")[1] == "echo" {
		param := strings.Split(splitedRequestLine[1], "/")[2]
		conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(param), param)))

	} else if strings.Split(splitedRequestLine[1], "/")[1] == "user-agent" {
		param := strings.Split(splitedUserAgent[1], ":")[0]
		conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(param), param)))
	} else {
		conn.Write([]byte(status404))
		return
	}

}
