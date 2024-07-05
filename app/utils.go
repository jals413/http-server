package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func requestHandler(conn net.Conn) {

	defer conn.Close()

	statusOK := "HTTP/1.1 200 OK\r\n\r\n"
	status404 := "HTTP/1.1 404 Not Found\r\n\r\n"
	statusCreated := "HTTP/1.1 201 Created\r\n\r\n"

	buffer := make([]byte, 1024)
	conn.Read(buffer)

	splitedHeader := strings.Split(string(buffer), "\r\n")
	splitedRequestLine := strings.Split(splitedHeader[0], " ")
	splitedUserAgent := strings.Split(splitedHeader[2], " ")

	var splitedAcceptEncoding []string
	encoded := false
	for _, header := range splitedHeader {
		if strings.HasPrefix(header, "Accept-Encoding") {
			splitedAcceptEncoding = strings.Split(header[len("Accept-Encoding: "):], ",")
			break

		}
	}
	for _, encoding := range splitedAcceptEncoding {
		if strings.TrimSpace(encoding) == "gzip" {
			encoded = true
			break
		}
	}

	if encoded {
		if splitedRequestLine[0] == "GET" {

			if splitedRequestLine[1] == "/" {
				conn.Write([]byte(statusOK))
				return
			} else if strings.Split(splitedRequestLine[1], "/")[1] == "echo" {
				param := strings.Split(splitedRequestLine[1], "/")[2]
				conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Encoding: gzip\r\nContent-Length: %d\r\n\r\n%s", len(param), param)))
				return

			} else if strings.Split(splitedRequestLine[1], "/")[1] == "user-agent" {
				param := strings.Split(splitedUserAgent[1], ":")[0]
				conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Encoding: gzip\r\nContent-Length: %d\r\n\r\n%s", len(param), param)))
				return

			} else if strings.Split(splitedRequestLine[1], "/")[1] == "files" {
				dir := os.Args[2]
				fileName := strings.Split(splitedRequestLine[1], "/")[2]
				data, err := os.ReadFile(dir + fileName)
				if err != nil {
					conn.Write([]byte(status404))
				} else {
					conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type:application/octet-stream\r\nContent-Encoding: gzip\r\nContent-Length: %d\r\n\r\n%s", len(data), data)))
				}
				return

			} else {
				conn.Write([]byte(status404))
				return
			}
		} else if splitedRequestLine[0] == "POST" {
			if strings.Split(splitedRequestLine[1], "/")[1] == "files" {

				dir := os.Args[2]
				fileName := strings.Split(splitedRequestLine[1], "/")[2]
				path := dir + fileName
				filecontent := []byte(strings.Trim(splitedHeader[len(splitedHeader)-1], "\x00"))

				fmt.Println(filecontent)
				fmt.Println(path)

				_ = os.WriteFile(path, []byte(filecontent), 0644)
				conn.Write([]byte(statusCreated))
				return
			} else {
				conn.Write([]byte(status404))
				return

			}
		}
	} else {
		if splitedRequestLine[0] == "GET" {

			if splitedRequestLine[1] == "/" {
				conn.Write([]byte(statusOK))
				return
			} else if strings.Split(splitedRequestLine[1], "/")[1] == "echo" {
				param := strings.Split(splitedRequestLine[1], "/")[2]
				conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(param), param)))
				return

			} else if strings.Split(splitedRequestLine[1], "/")[1] == "user-agent" {
				param := strings.Split(splitedUserAgent[1], ":")[0]
				conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(param), param)))
				return

			} else if strings.Split(splitedRequestLine[1], "/")[1] == "files" {
				dir := os.Args[2]
				fileName := strings.Split(splitedRequestLine[1], "/")[2]
				data, err := os.ReadFile(dir + fileName)
				if err != nil {
					conn.Write([]byte(status404))
				} else {
					conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type:application/octet-stream\r\nContent-Length: %d\r\n\r\n%s", len(data), data)))
				}
				return

			} else {
				conn.Write([]byte(status404))
				return
			}
		} else if splitedRequestLine[0] == "POST" {
			if strings.Split(splitedRequestLine[1], "/")[1] == "files" {

				dir := os.Args[2]
				fileName := strings.Split(splitedRequestLine[1], "/")[2]
				path := dir + fileName
				filecontent := []byte(strings.Trim(splitedHeader[len(splitedHeader)-1], "\x00"))

				fmt.Println(filecontent)
				fmt.Println(path)

				_ = os.WriteFile(path, []byte(filecontent), 0644)
				conn.Write([]byte(statusCreated))
				return
			} else {
				conn.Write([]byte(status404))
				return

			}
		}
	}

}
