package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	addr := "localhost:8000"

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("listening on addr: %s\n", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}

		log.Printf("accept one connection: %s\n", conn.RemoteAddr())
		go handleConn(conn) // handle connections concurrently
	}
}

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		log.Println("wg done 1")
	}()

	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)

	var wg sync.WaitGroup
	for input.Scan() {
		wg.Add(1)
		log.Println("wg add 1")

		go echo(c, input.Text(), 1*time.Second, &wg) //struct值拷贝，只能传指针
	}

	log.Printf("input scan finished, wg(%v) waiting\n", wg)
	wg.Wait()
	log.Println("wg wait passed")

	// NOTE: ignoring potential errors from input.Err()
	c.Close()
	log.Println("conn closed")
}
