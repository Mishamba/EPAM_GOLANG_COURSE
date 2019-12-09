package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

func handleAns(conn net.Conn) {
	defer fmt.Println(conn.RemoteAddr().String(), " closed")
	defer conn.Close()
	for {
		source, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("can't get message from client")
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		fmt.Println("received ", source)

		str := strings.TrimSuffix(source, string('\n'))
		if data, err := strconv.Atoi(strings.TrimSuffix(str, string('\n'))); err == nil {
			_, err := conn.Write([]byte(string(data * 2)))
			if err != nil {
				fmt.Println("can't send answer for ", conn.RemoteAddr().String())
				fmt.Println(err)
				continue
			}
		} else {
			_, err := conn.Write([]byte(strings.ToUpper(string(data))))
			if err != nil {
				fmt.Println("can't send answer for ", conn.RemoteAddr().String())
				fmt.Println(err)
				continue
			}
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8083")
	defer ln.Close()
	if err != nil {
		log.Fatal("1) ", err)
	}

	for {
		conn, err := ln.Accept()
		fmt.Println(conn.RemoteAddr().String(), " connected")
		if err != nil {
			fmt.Print("2) ")
			panic(err)
		}
		go handleAns(conn)
	}
}
