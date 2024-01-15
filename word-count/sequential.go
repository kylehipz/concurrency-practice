package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

func ff(r rune) bool {
	return !unicode.IsLetter(r)
}

func readFile(file string) map[string]int {
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
	return wm
}

func RunSequential() {
	root := os.Args[1]

	filenames := getFiles(root)

	wordmap := make(map[string]int)

	for _, filename := range filenames {
		wm := readFile(filename)
		for w, c := range wm {
			wordmap[w] += c
		}
	}

	outFile, _ := os.Create("out/out.txt")
	defer outFile.Close()

	for w, c := range wordmap {
		fmt.Fprintf(outFile, "%v %v\n", w, c)
	}
}
