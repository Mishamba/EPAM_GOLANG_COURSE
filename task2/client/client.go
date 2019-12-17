package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", ":8083")
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to ", conn.RemoteAddr().String())
	defer fmt.Println("disconnected from ", conn.RemoteAddr().String())
	defer conn.Close()
	for {
		input := bufio.NewReader(os.Stdin)
		fmt.Println("enter data to send")
		data, err := input.ReadString('\n')
		if err != nil {
			break
		}
		check := strings.TrimSuffix(data, string('\n'))
		if check == "exit" {
			break
		}
		conn.Write([]byte(data + string('\n')))
		ans, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(ans)
		}
	}
}
