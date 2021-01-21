package main

import (
	"log"
	"os"

	"github.com/DaneSpiritGOD/ex8.6/links"
)

func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	dm := &depthManager{3, 1, int32(len(os.Args[1:]))}

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				log.Printf("dm before crawl: %v\n", dm)
				foundLinks := crawl(link)
				go func() {
					worklist <- foundLinks
					dm.addWorks()
					log.Printf("dm after add work: %v\n", dm)
				}()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}

		dm.removeWorks()
		log.Printf("dm after remove work: %v\n", dm)

		if dm.canLeave() {
			log.Println("break")
			break
		}

		if dm.canIncreaseDepth() {
			dm.increaseDepth()
			log.Printf("dm after increate depth: %v\n", dm)
		}
	}
}

func crawl(url string) []string {
	log.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
