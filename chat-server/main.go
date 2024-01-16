package main

import (
	"bufio"
	"fmt"
	"net"
)

type Client = chan string

func broadcast(messages Client, activeClients, leavingClients chan Client) {
	clients := make(map[Client]bool)
	for {
		select {
		case msg := <-messages:
			for client := range clients {
				client <- msg
			}
		case client := <-activeClients:
			clients[client] = true
		case client := <-leavingClients:
			delete(clients, client)
			close(client)
		}
	}
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening on tcp port 8000")

	messages := make(Client)
	activeClients := make(chan Client)
	leavingClients := make(chan Client)

	go broadcast(messages, activeClients, leavingClients)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConn(conn, messages, activeClients, leavingClients)
	}
}

func handleConn(conn net.Conn, messages Client, activeClients, leavingClients chan Client) {
	defer conn.Close()
	connCh := make(chan string)
	go func() {
		for msg := range connCh {
			fmt.Fprintln(conn, msg)
		}
	}()

	addr := conn.RemoteAddr().String()
	connCh <- "Welcome " + addr + "!"
	messages <- addr + " has joined"
	activeClients <- connCh

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		message := scanner.Text()

		messages <- addr + ": " + message
	}

	messages <- addr + " has left"
	leavingClients <- connCh
}
