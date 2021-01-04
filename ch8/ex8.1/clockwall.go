// Netcat1 is a read-only TCP client.
package main

import (
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// .\clockwall ShangHai=localhost:8081 Chicago=localhost:8082 Accra=localhost:8083
func main() {
	args := os.Args[1:]

	for _, pair := range args {
		items := strings.Split(pair, "=")
		if len(items) != 2 {
			log.Fatal("the format of input is wrong")
		}

		go func(name string, addr string) {
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("connection is successful. name: %s addr: %s", name, addr)

			defer conn.Close()

			mustCopy(os.Stdout, conn)
		}(items[0], items[1])
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
