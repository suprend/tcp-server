package main

import (
	"bufio"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	reply, err := reader.ReadString('\n')
	if err != nil || reply != "OK\n" {
		panic("Error")
	}
}
