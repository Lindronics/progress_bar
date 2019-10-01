package main

import (
	"fmt"
	"strings"
	"time"
)

// Bar is a progress bar
type Bar struct {
	length int
	maxVal int
	t0     time.Time
}

func (b *Bar) update(i int) {
	percentage := 100 * float32(i) / float32(b.maxVal)
	level := b.length * i / b.maxVal
	progress := strings.Repeat("▓", level)
	blanks := strings.Repeat(" ", b.length-level)
	elapsed := time.Since(b.t0).Seconds()
	fmt.Printf("\rProgress: │%s%s│ %.2f%% - %.2fs", progress, blanks, percentage, elapsed)
}

// New creates a Bar tracking progress from 0 to maxVal
func New(maxVal int) Bar {
	return Bar{length: 50, maxVal: maxVal}
}

// Start sets up a progress bar
func (b *Bar) Start() {
	b.t0 = time.Now()
	fmt.Printf("\n")
	b.update(0)
}

// Set sets a new value of the progress bar
func (b *Bar) Set(i int) {
	b.update(i)
}

// Finish ends a progress bar
func (b *Bar) Finish() {
	b.update(b.maxVal)
	elapsed := time.Since(b.t0)
	fmt.Printf("\nWall time: %f\n", elapsed.Seconds())
}

func main() {
	b := New(100)
	b.Start()
	for i := 0; i < 100; i++ {
		b.Set(i)
		time.Sleep(30 * time.Millisecond)
	}
	b.Finish()
}
