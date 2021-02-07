// Package bank provides a concurrency-safe bank with one account.
package main

import "fmt"

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func deposit(amount int) { deposits <- amount }

func balance() int { return <-balances }

type callback struct {
	balance int
	error
}

type withDrawsI struct {
	amount int
	c      chan *callback
}

var withDraws = make(chan *withDrawsI)

func drawWith(amount int) (int, error) {
	c := make(chan *callback)
	withDraws <- &withDrawsI{amount, c}
	b := <-c
	return b.balance, b.error
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case w := <-withDraws:
			amount := w.amount
			if balance < amount {
				w.c <- &callback{balance, fmt.Errorf("balance(%d) is low than amount(%d)", balance, amount)}
			} else {
				balance -= amount
				w.c <- &callback{balance, nil}
			}
		case balances <- balance:
		}
	}
}

func init1() {
	go teller() // start the monitor goroutine
}

func main() {
	init1()

	deposit(200)
	deposit(100)

	drawWith2 := func(amount int) {
		balance, err := drawWith(amount)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("current balance: %d.\n", balance)
		}
	}

	drawWith2(700)
	drawWith2(100)
	drawWith2(200)
	drawWith2(100)
}
