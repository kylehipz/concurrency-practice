package main

import (
	"fmt"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	fmt.Println("Listening on tcp port 8000")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		defer conn.Close()

		for {
			buf := make([]byte, 1024)
			_, err = conn.Read(buf)

			fmt.Println(string(buf))
		}
	}
}
