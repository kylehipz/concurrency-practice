package main

import (
	"fmt"
	"time"
)

func main() {
	started := time.Now()

	RunSequential()

	elapsed := time.Since(started)
	fmt.Printf("Finished: %.2fs\n", elapsed.Seconds())
}
