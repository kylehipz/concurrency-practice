package main

import (
	"log"
	"os"
	"path/filepath"
)

func getFiles(dir string) []string {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalln(err)
	}

	files := []string{}

	for _, entry := range entries {
		path := filepath.Join(dir, entry.Name())
		files = append(files, path)
	}

	return files
}
