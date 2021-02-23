package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

type sss struct{}

func main() {
	ping := make(chan sss)
	pong := make(chan sss)

	i := 0

	start := time.Now()
	go func() {
		for {
			i++
			pong <- <-ping
		}
	}()

	go func() {
		ping <- sss{}

		for {
			ping <- <-pong
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	consume := time.Since(start)
	roundTripPerSecond := float64(i) / consume.Seconds()
	fmt.Printf("%f round trips per second", roundTripPerSecond)
}
