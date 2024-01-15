package main

import (
	"os"
	"path/filepath"
	"sync"
)

func walkDirCon(dir string, wg *sync.WaitGroup, fileSizes chan int64) {
	defer wg.Done()

	for _, entry := range getDirEntries(dir) {
		filename := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			wg.Add(1)
			go walkDirCon(filename, wg, fileSizes)
		} else {
			info, _ := entry.Info()
			fileSizes <- info.Size()
		}
	}
}

func RunConcurrent() {
	root := os.Args[1]
	count := 0
	size := int64(0)

	fileSizes := make(chan int64)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go walkDirCon(root, &wg, fileSizes)

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	for sz := range fileSizes {
		count++
		size += sz
	}

	printDiskUsage(count, size)
}
