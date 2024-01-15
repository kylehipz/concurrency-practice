package main

import (
	"os"
	"path/filepath"
)

func walkDirSeq(dir string, count *int, size *int64) {
	for _, entry := range getDirEntries(dir) {
		if entry.IsDir() {
			path := filepath.Join(dir, entry.Name())
			walkDirSeq(path, count, size)
		} else {
			info, _ := entry.Info()

			*count++
			*size += info.Size()
		}
	}
}

func RunSequential() {
	root := os.Args[1]
	count := 0
	size := int64(0)

	walkDirSeq(root, &count, &size)

	printDiskUsage(count, size)
}
