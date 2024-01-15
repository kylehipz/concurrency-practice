package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func getDirEntries(dir string) []fs.DirEntry {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	return entries
}

func getDirEntriesSema(dir string, sema chan bool) []fs.DirEntry {
	// Control concurrency by counting semaphores
	sema <- true
	defer func() { <-sema }()

	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	return entries
}

func printDiskUsage(count int, size int64) {
	fmt.Printf("Total files: %d, Total size: %.1f GB\n", count, float64(size)/1e9)
}
