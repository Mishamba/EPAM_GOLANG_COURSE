package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func handleAns(conn net.Conn) {
	defer fmt.Println(conn.RemoteAddr().String(), " closed")
	defer conn.Close()
	for {
		m, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}

		var resp string
		m = strings.TrimSuffix(m, string('\n'))

		n, err := strconv.Atoi(m)
		if err != nil {
			resp = strings.ToUpper(m)
		} else {
			n *= 2
			resp = strconv.Itoa(n)
		}

		_, err = conn.Write([]byte(resp + string('\n')))
		if err != nil {
			fmt.Println(err)
			break
		}

	}
}

func main() {
	ln, err := net.Listen("tcp", ":8083")
	defer ln.Close()
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		fmt.Println(conn.RemoteAddr().String(), " connected")
		if err != nil {
			panic(err)
		}
		go handleAns(conn)
	}
}
