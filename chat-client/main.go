package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)

	go func() {
		connSc := bufio.NewScanner(conn)

		for connSc.Scan() {
			message := connSc.Text()

			fmt.Println(message)
		}
	}()

	for scanner.Scan() {
		message := scanner.Text()

		if message != "" {
			fmt.Fprintln(conn, message)
		}
	}
}
