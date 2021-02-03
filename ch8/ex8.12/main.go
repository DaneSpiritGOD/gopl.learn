package main

import (
	"bufio"
	"fmt"
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

}

func broadcaster() {
	clients := make(map[client]bool)

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
			} else {
				return
			}
		case cli, ok := <-leaving:
			if ok {
				delete(clients, cli)
				close(cli.ear)

				for c := range clients {

				}
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
