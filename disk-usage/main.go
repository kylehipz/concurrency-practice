package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	RunConcurrent()
	elapsed := time.Since(start)

	fmt.Printf("Finished: %.2fs\n", elapsed.Seconds())
}
