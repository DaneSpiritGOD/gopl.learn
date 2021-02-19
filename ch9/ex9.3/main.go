package main

import (
	"fmt"
	"log"
	"time"

	"github.com/DaneSpiritGOD/ex9.3/memo"
)

func main() {
	f := func(key string, done <-chan struct{}) (interface{}, error) {
		select {
		case <-done:
			return nil, memo.ErrorDone
		case <-time.After(3 * time.Second):
			if key == "" {
				return nil, fmt.Errorf("key cannot be empty")
			}

			return key + "a", nil
		}
	}

	doneKey := "#"
	keys := []string{"a", "a", "b", "c", "c", doneKey}
	m := memo.New(f)

	for _, key := range keys {

		calcFunc := func(s string, d chan struct{}) {
			log.Printf("caculating result of key: %s", key)
			res, err := m.Get(s, d)
			if err != nil {
				switch err {
				case memo.ErrorDone:
					log.Printf("done interrupt when calculating result of key: %s", s)
				default:
					log.Print(err)
				}
			} else {
				log.Printf("result: %v", res)
			}
		}

		const tryCount = 3
		if key == doneKey {
			for index := range [tryCount]int{} {
				if index == 0 {
					done := make(chan struct{})
					go func() {
						time.AfterFunc(2*time.Second, func() {
							done <- struct{}{}
						})
					}()

					calcFunc(key, done)
				} else {
					done := make(chan struct{})
					calcFunc(key, done)
				}
			}

		} else {
			done := make(chan struct{})
			calcFunc(key, done)
		}
	}
}
