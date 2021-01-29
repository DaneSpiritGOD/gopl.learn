package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"sync/atomic"
	"time"

	"github.com/DaneSpiritGOD/ex8.7/links"
)

// go build -o main.exe main.go types.go
// .\main.exe https://books.studygolang.com/gopl-zh/ ./html_files *>&1 > main.log &
func main() {
	worklist := make(chan worksList) // lists of URLs, may have duplicates
	unseenWork := make(chan works)   // de-duplicated URLs

	home := os.Args[1]
	homeURL, err := url.Parse(home)
	if err != nil {
		log.Fatalf("%s is not a valid url", home)
	}

	host := homeURL.Host

	homeStoreDir, err := filepath.Abs(os.Args[2])
	if err != nil {
		log.Fatalf("%s is not a valid directory!", homeStoreDir)
	}

	log.Printf("save directory: %s", homeStoreDir)
	if _, err := os.Stat(homeStoreDir); os.IsNotExist(err) {
		log.Printf("directory: %s is not exists, need to create", homeStoreDir)
		os.Mkdir(homeStoreDir, 0)
	} else {
		os.RemoveAll(homeStoreDir)
		os.Mkdir(homeStoreDir, 0)
	}

	go func() { worklist <- worksList{[]string{home}, 0} }()

	fileIndex := int32(0)
	pFileIndex := &fileIndex

	getStorePathDunc := func() string {
		atomic.AddInt32(pFileIndex, 1)
		result := filepath.Join(homeStoreDir, fmt.Sprintf("%d.html", atomic.LoadInt32(pFileIndex)))
		return result
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

						linkURL, _ := url.Parse(link)
						linkHost := linkURL.Host
						if linkHost != host {
							log.Printf("the host %s of url %s is not equals to %s", linkHost, link, host)
							continue
						}

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
