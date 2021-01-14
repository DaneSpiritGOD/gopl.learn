package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	done := make(chan struct{})
	go func() {
		if _, err := io.Copy(os.Stdout, conn); err != nil {
			log.Println(err)
		}

		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()

	mustCopy(conn, os.Stdin)

	//Ctrl+Z
	log.Println("os.Stdin is closed by Ctrl+Z, next step is closewrite tcpConn")
	tcpConn := conn.(*net.TCPConn)
	tcpConn.CloseWrite()

	log.Println("waiting for ch named 'done'")
	<-done // wait for background goroutine to finish
	log.Println("ch named 'done' is done")
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil && err != io.EOF {
		log.Fatal(err)
	}
}
