package main

import (
	"bufio"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		go func(c net.Conn) {
			defer c.Close()
			w := bufio.NewWriter(c)
			w.WriteString("OK\n")
			w.Flush()
		}(conn)
	}
}
