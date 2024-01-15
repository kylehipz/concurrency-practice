package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	RunSequential()
	elapsed := time.Since(start)

	fmt.Printf("Finished: %s\n", elapsed)
}
