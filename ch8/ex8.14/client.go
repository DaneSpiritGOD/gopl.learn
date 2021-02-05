package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	name := os.Args[1]

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Print(err)
		return
	}

	defer conn.Close()

	io.WriteString(conn, name+"\n")

	done := make(chan struct{})
	go func() {
		mustCopy(os.Stdout, conn)
		log.Printf("remote is closed")

		done <- struct{}{}
	}()

	go func() {
		mustCopy(conn, os.Stdin)
		log.Printf("stdin is closed")

		done <- struct{}{}
	}()

	<-done
	log.Printf("quiting")
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
