package main

import (
	"fmt"
	"path/filepath"

	"github.com/google/uuid"
)

func walkDir(dir string) {
	for _, entry := range getDirEntries(dir) {
		name := entry.Name()
		fileOrDir := filepath.Join(dir, name)

		if entry.IsDir() {
			filename := filepath.Join(dir, name)
			walkDir(filename)
		} else {
			newFilename := fmt.Sprintf("out/%s-%s", name, uuid.NewString())
			copyFile(fileOrDir, newFilename)
		}
	}
}

func RunSequential() {
	walkDir("in")
}
