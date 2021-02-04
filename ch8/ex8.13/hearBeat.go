package main

import (
	"net"
	"time"
)

type heartBeat struct {
	expire time.Duration
	conn   net.Conn
	timer  *time.Timer
}

func newBeat(expire time.Duration, conn net.Conn) *heartBeat {
	return &heartBeat{expire, conn, nil}
}
