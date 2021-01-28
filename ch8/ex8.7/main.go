package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/DaneSpiritGOD/ex8.7/links"
)

// go build -o main.exe main.go types.go
// .\main.exe https://books.studygolang.com/gopl-zh/ ./html_files *>&1 > main.log
func main() {
	worklist := make(chan worksList) // lists of URLs, may have duplicates
	unseenWork := make(chan works)   // de-duplicated URLs

	homeURL := []string{os.Args[1]}

	go func() { worklist <- worksList{homeURL, 0} }()

	homeStoreDir, err := filepath.Abs(os.Args[2])
	if err != nil {
		log.Fatalf("%s is not a valid directory!", homeStoreDir)
	}

	log.Printf("save directory: %s", homeStoreDir)
	if _, err := os.Stat(homeStoreDir); os.IsNotExist(err) {
		log.Printf("directory: %s is not exists, need to create", homeStoreDir)
		os.Mkdir(homeStoreDir, 0)
	}

	fileIndex := 1
	getStorePathDunc := func() string {
		return filepath.Join(homeStoreDir, fmt.Sprintf("%d.html", fileIndex))
	}

	const maxDepth = 5000 // 最大深度

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for work := range unseenWork {
				depth := work.depth
				url := work.link
				fn := getStorePathDunc()

				log.Printf("current depth: %d, url: %s, fileName: %s", depth, url, fn)
				foundLinks := crawl(url, fn)

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

func crawl(url string, filePath string) []string {
	list, err := links.SaveAndExtract(url, filePath)
	if err != nil {
		log.Print(err)
	}
	return list
}
