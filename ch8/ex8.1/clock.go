// Clock1 is a TCP server that periodically writes the time.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var port = flag.Int("port", 8090, "port to listen")
var zone = flag.String("zone", "China/BeiJing", "location")

// .\clock -port=8081 -zone=Asia/Shanghai *>&1 > as.txt &
// .\clock -port=8082 -zone=America/Chicago *>&1 > ac.txt &
// .\clock -port=8083 -zone=Africa/Accra *>&1 > aa.txt &
func main() {
	flag.Parse()

	addr := fmt.Sprintf("localhost:%d", *port)

	name := *zone
	location, err := time.LoadLocation(name)
	if err != nil {
		log.Fatalf("load location error: %v\n", err)
	}
	log.Printf("location: %s\n", location)

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
		go handleConn(conn, name, location) // handle connections concurrently
	}
}

func handleConn(c net.Conn, zoneName string, location *time.Location) {
	defer c.Close()

	for {
		t := fmt.Sprintf("%s: %s", zoneName, time.Now().In(location).Format("15:04:05\n"))

		_, err := io.WriteString(c, t)
		if err != nil {
			log.Printf("net error: %v\n", err)
			return
		}

		time.Sleep(1 * time.Second)
	}
}
