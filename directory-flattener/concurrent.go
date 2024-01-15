package main

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/google/uuid"
)

type FileCP struct {
	src string
	dst string
}

func walkDirCon(dir string, wg *sync.WaitGroup, results chan<- FileCP) {
	defer wg.Done()

	for _, entry := range getDirEntries(dir) {
		name := entry.Name()
		fileOrDir := filepath.Join(dir, name)

		if entry.IsDir() {
			filename := filepath.Join(dir, name)
			wg.Add(1)
			walkDirCon(filename, wg, results)
		} else {
			newFilename := fmt.Sprintf("out/%s-%s.txt", name, uuid.NewString())
			results <- FileCP{fileOrDir, newFilename}
		}
	}
}

func processResults(results <-chan FileCP) {
	for res := range results {
		copyFile(res.src, res.dst)
	}
}

func RunConcurrent() {
	dir := "in"
	numWorkers := 30
	wg := sync.WaitGroup{}
	results := make(chan FileCP)
	done := make(chan bool)

	wg.Add(1)
	go walkDirCon(dir, &wg, results)

	go func() {
		wg.Wait()
		close(results)
		done <- true
		close(done)
		fmt.Println("Done")
	}()

	for i := 0; i < numWorkers; i++ {
		go processResults(results)
	}

	<-done
}
