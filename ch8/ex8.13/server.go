package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

type client struct {
	name string
	ear  chan<- string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

var logger = log.New

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}

func broadcaster() {
	clients := make(map[client]bool)

	getAllClientsName := func() string {
		buf := &bytes.Buffer{}

		io.WriteString(buf, "online clients:")
		for c := range clients {
			io.WriteString(buf, " ")
			io.WriteString(buf, c.name)
		}

		return buf.String()
	}

	for {
		select {
		case msg, ok := <-messages:
			if ok {
				for cli := range clients {
					cli.ear <- msg
				}
			} else {
				return
			}
		case cli, ok := <-entering:
			if ok {
				clients[cli] = true

				go func() {
					messages <- cli.name + " enters"
					messages <- getAllClientsName()
				}()
			} else {
				return
			}
		case cli, ok := <-leaving:
			if ok {
				delete(clients, cli)
				close(cli.ear)

				go func() {
					messages <- cli.name + " leaves"
					messages <- getAllClientsName()
				}()
			} else {
				return
			}
		}
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	ear := make(chan string)
	go sendEarContent(conn, ear) // listen with ear, and send message to remote

	name := conn.RemoteAddr().String()
	cli := client{name, ear}

	entering <- cli // add to entering

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- name + ": " + input.Text()
	}

	leaving <- cli
}

func sendEarContent(conn net.Conn, ear <-chan string) {
	for msg := range ear {
		fmt.Fprintln(conn, msg)
	}
}
