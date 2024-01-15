package main

import (
	"os"
	"path/filepath"
	"sync"
	"time"
)

func walkDirCon(dir string, wg *sync.WaitGroup, fileSizes chan int64, sema chan bool) {
	defer wg.Done()

	for _, entry := range getDirEntriesSema(dir, sema) {
		filename := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			wg.Add(1)
			go walkDirCon(filename, wg, fileSizes, sema)
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
	sema := make(chan bool, 50)
	wg := sync.WaitGroup{}

	// Walk directory
	wg.Add(1)
	go walkDirCon(root, &wg, fileSizes, sema)

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	tick := time.Tick(500 * time.Millisecond)

loop:
	for {
		select {
		case sz, ok := <-fileSizes:
			if !ok {
				break loop
			}
			count++
			size += sz
		case <-tick:
			printDiskUsage(count, size)
		}
	}

	printDiskUsage(count, size)
}
