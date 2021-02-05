package main

import (
	"log"
	"net"
	"time"
)

type heartBeat struct {
	expire time.Duration
	conn   *net.Conn
	timer  *time.Timer
}

type client struct {
	name string
	ear  chan<- string
}

func newBeat(expire time.Duration, conn *net.Conn) *heartBeat {
	return &heartBeat{expire, conn, nil}
}

func (hb *heartBeat) start() {
	hb.resetTimer()
}

func (hb *heartBeat) resetTimer() {
	hb.timer = hb.createTimer()
}

func (hb *heartBeat) createTimer() *time.Timer {
	return time.AfterFunc(hb.expire, func() {
		log.Println("time out")
		(*hb.conn).Close()
	})
}

func (hb *heartBeat) reset() bool {
	if !hb.timer.Stop() { // timer fired already
		return false // reset failed
	}

	hb.resetTimer()
	return true
}
