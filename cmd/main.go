package main

import (
	"time"
	pbar "github.com/philhanna/progress-bar"
)

const limit = 25000

// Demonstration of the progress bar
func main() {
	pbar := pbar.NewProgress(limit)
	for i := 0; i < limit; i++ {
		time.Sleep(time.Millisecond)
		pbar.Add(1)
	}
}