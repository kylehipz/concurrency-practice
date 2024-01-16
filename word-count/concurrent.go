package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readFileCon(file string, results chan<- map[string]int) {
	wm := make(map[string]int)

	f, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln(err)
		}

		words := strings.FieldsFunc(string(line), ff)

		for _, w := range words {
			wm[w]++
		}
	}

	fmt.Printf("Finished processing file: %s\n", file)
	results <- wm
}

func processFiles(files <-chan string, results chan<- map[string]int) {
	for file := range files {
		readFileCon(file, results)
	}
}

func RunConcurrent() {
	root := os.Args[1]
	filenames := getFiles(root)
	fileCount := len(filenames)

	wordmap := make(map[string]int)

	numWorkers := 20
	files := make(chan string, fileCount)
	results := make(chan map[string]int, fileCount)

	for i := 0; i < numWorkers; i++ {
		go processFiles(files, results)
	}

	go func() {
		for _, file := range filenames {
			files <- file
		}
	}()

	for i := 0; i < fileCount; i++ {
		wm := <-results

		for w, c := range wm {
			wordmap[w] += c
		}
	}

	outFile, _ := os.Create("out/out2.txt")
	defer outFile.Close()

	for w, c := range wordmap {
		fmt.Fprintf(outFile, "%v %v\n", w, c)
	}
}
