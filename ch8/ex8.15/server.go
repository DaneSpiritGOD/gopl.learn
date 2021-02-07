package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string, 20)
)

var timeout = flag.Int("t", 20, "timeout of per connection, unit: second")

// go build -o server.exe server.go hearBeat.go
// go build -o client.exe client.go
// .\server.exe -t 20
func main() {
	flag.Parse()
	seconds := *timeout

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

		go handleConn(conn, seconds)
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
			if !ok {
				return
			}

			for cli := range clients {
				select {
				case cli.ear <- msg:
				case <-time.After(1 * time.Second):
					log.Printf("skip one message which is sent to client: %s", cli.name)
				}
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

func handleConn(conn net.Conn, seconds int) {
	defer conn.Close()

	heartBeat := newBeat(time.Duration(seconds)*time.Second, conn)
	heartBeat.start()

	var name string

	input := bufio.NewScanner(conn)
	if input.Scan() {
		if err := input.Err(); err != nil {
			log.Print(err)
			return
		}

		name = input.Text()
	}

	ear := make(chan string, 20) // channel which has buffer
	go sendEarContent(conn, ear) // listen with ear, and send message to remote

	cli := client{name, ear}
	entering <- cli // add to entering

	getKickSay := func() {
		messages <- fmt.Sprintf("%s didn't speak anything in some time and is kicked out", name)
	}

	for input.Scan() {
		if !heartBeat.reset() {
			getKickSay()
			break // reset failed
		}

		messages <- name + ": " + input.Text()
	}

	if err := input.Err(); err != nil {
		go getKickSay()
	}

	leaving <- cli
}

func sendEarContent(conn net.Conn, ear <-chan string) {
	for msg := range ear {
		fmt.Fprintln(conn, msg)
	}
}
