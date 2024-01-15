package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	elapsed := time.Since(start)

	fmt.Printf("Finished: %s\n", elapsed)
}
