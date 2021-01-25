package main

import (
	"log"
	"os"
	"time"

	"github.com/DaneSpiritGOD/ex8.6/links"
)

func main() {
	worklist := make(chan worksList) // lists of URLs, may have duplicates
	unseenWork := make(chan works)   // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() { worklist <- worksList{os.Args[1:], 0} }()

	const maxDepth = 1 // 最大深度

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for work := range unseenWork {
				depth := work.depth
				url := work.link
				log.Printf("current depth: %d, url: %s", depth, url)
				foundLinks := crawl(url)

				if depth == maxDepth {
					log.Println("reach max depth")
				} else {
					go func(depth int) {
						worklist <- worksList{foundLinks, depth + 1}
					}(depth)
				}
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)

forloop:
	for {
		select {
		case list, ok := <-worklist:
			{
				if ok {
					for _, link := range list.links {
						if !seen[link] {
							seen[link] = true
							unseenWork <- works{link, list.depth}
						}
					}
				} else {
					break forloop
				}
			}
		case <-time.After(10 * time.Second):
			log.Println("time exceed")

			close(unseenWork)
			close(worklist)
			break forloop
		}
	}

	log.Println("main exiting")
}

func crawl(url string) []string {
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
