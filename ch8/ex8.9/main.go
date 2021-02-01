package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var directoryF = flag.String("d", ".", "the directory to monitor")
var periodF = flag.Int("p", 5, "period(seconds) of showing size of directory")

// go run main.go -d "D:\360极速浏览器下载" -p 3
func main() {
	flag.Parse()

	dir := *directoryF
	period := *periodF
	log.Printf("dir: %s period: %ds\n", dir, period)

	tick := time.Tick(time.Duration(period) * time.Second)
	for {
		select {
		case <-tick:
			fileSizes := make(chan int64)

			wg := &sync.WaitGroup{}
			wg.Add(1)
			go func() {
				walkDir(dir, fileSizes, wg)
			}()

			go func() {
				wg.Wait()
				close(fileSizes)
			}()

			var nbytes int64
			var nfiles int64
			for size := range fileSizes {
				nfiles++
				nbytes += size
			}

			printDiskUsage(nfiles, nbytes)
		}
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	log.Printf("| %d files  %.3f GB\n", nfiles, float64(nbytes)/1e9)
}

var sema = make(chan struct{}, 20)

func getDirEntries(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("read dir error: %v\n", err)
	}
	return entries
}

func walkDir(dir string, fileSizes chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()

	entries := getDirEntries(dir)
	for _, e := range entries {
		if e.IsDir() {
			wg.Add(1)
			subDir := filepath.Join(dir, e.Name())

			go func() {
				walkDir(subDir, fileSizes, wg)
			}()
		} else {
			fileSizes <- e.Size()
		}
	}
}
