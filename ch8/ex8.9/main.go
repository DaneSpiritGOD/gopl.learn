package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var directoryF = flag.String("d", ".", "the directory to monitor")
var periodF = flag.Int("p", 5, "period of showing size of directory")

func main() {
	flag.Parse()

	dir := *directoryF
	period := *periodF

	log.Printf("dir: %s, period: %d\n", dir, period)
}

func getDirEntries(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("read dir error: %v\n", err)
	}
	return entries
}

func walkDir(dir string, fileSizes chan<- int64) {
	entries := getDirEntries(dir)
	for _, e := range entries {
		if e.IsDir() {
			subDir := filepath.Join(dir, e.Name())
			walkDir(subDir, fileSizes)
		} else {
			fileSizes <- e.Size()
		}
	}
}
