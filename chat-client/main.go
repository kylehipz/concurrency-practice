package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	for {
		fmt.Println("Hey")
		_, err = conn.Write([]byte("Hello"))

		time.Sleep(500 * time.Millisecond)

		if err != nil {
			panic(err)
		}
	}
}
