package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func walkDir(dir string, count *int, size *int64) {
	for _, entry := range getDirEntries(dir) {
		if entry.IsDir() {
			path := filepath.Join(dir, entry.Name())
			walkDir(path, count, size)
		} else {
			info, _ := entry.Info()

			*count++
			*size += info.Size()
		}
	}
}

func getDirEntries(dir string) []fs.DirEntry {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	return entries
}

func RunSequential() {
	root := os.Args[1]
	count := 0
	size := int64(0)

	walkDir(root, &count, &size)

	printDiskUsage(count, size)
}

func printDiskUsage(count int, size int64) {
	fmt.Printf("Total files: %d, Total size: %.1f GB\n", count, float64(size)/1e9)
}
