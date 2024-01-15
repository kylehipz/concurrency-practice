package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
)

func getDirEntries(dir string) []fs.DirEntry {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalln(err)
	}

	return entries
}

func copyFile(srcPath, dstPath string) {
	fmt.Printf("Copying %s to %s\n", srcPath, dstPath)
	src, err := os.Open(srcPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)

	if err != nil {
		log.Fatalln(err)
	}
}
