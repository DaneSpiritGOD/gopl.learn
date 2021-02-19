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
			return nil, fmt.Errorf("done interrupt when caculating result of key: %s", key)
		case <-time.After(3 * time.Second):
			if key == "" {
				return nil, fmt.Errorf("key cannot be empty")
			}

			return key + "a", nil
		}
	}

	keys := []string{"a", "a", "a", "b", "c", "c", " ", " ", " ", " "}
	m := memo.New(f)

	for _, key := range keys {
		done := make(chan struct{})

		if key == " " {
			go func() {
				time.AfterFunc(1*time.Second, func() {
					done <- struct{}{}
				})
			}()
		}

		log.Printf("caculating result of key: %s", key)

		res, err := m.Get(key, done)
		if err != nil {
			log.Print(err)
		} else {
			log.Printf("result: %v", res)
		}
	}
}
