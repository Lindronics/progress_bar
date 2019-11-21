// Package progressbar contains a very basic console progress bar for Go programs
// TODO: implement error handling
package progressbar

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

// Bar is a simple console progress bar.
type Bar struct {
	width          int
	val            int
	maxVal         int
	startTime      time.Time
	theme          Theme
	showPercentage bool
	showTime       bool
	isFinished     bool
	lock           sync.Mutex
}

// Theme determines the look of a Bar.
type Theme struct {
	StartChar    rune
	EndChar      rune
	ProgressChar rune
}

// -------------------
// Optional parameters
// -------------------

// BarWidth returns a function for setting the width of a Bar.
func BarWidth(width int) func(*Bar) error {
	return func(b *Bar) error {
		if width <= 0 {
			return errors.New("width must be positive")
		}
		b.width = width
		return nil
	}
}

// BarTheme returns a function for setting the visual theme of a Bar.
func BarTheme(theme Theme) func(*Bar) error {
	return func(b *Bar) error {
		if theme.StartChar == 0 || theme.EndChar == 0 || theme.ProgressChar == 0 {
			return errors.New("invalid theme characters")
		}
		b.theme = theme
		return nil
	}
}

// BarShowPercent returns a function for setting whether the Bar displays a percentage.
func BarShowPercent(showPercentage bool) func(*Bar) error {
	return func(b *Bar) error {
		b.showPercentage = showPercentage
		return nil
	}
}

// BarShowTime returns a function for setting whether the Bar displays elapsed time.
func BarShowTime(showTime bool) func(*Bar) error {
	return func(b *Bar) error {
		b.showTime = showTime
		return nil
	}
}

// -------------------
// Methods
// -------------------

// New creates a Bar tracking progress from 0 to maxVal.
// Takes optional arguments as varidadic functions.
// Returns a new Bar object.
func New(maxVal int, kwargs ...func(*Bar) error) (*Bar, error) {

	if maxVal <= 0 {
		return nil, errors.New("invalid maxVal")
	}

	theme := Theme{
		StartChar:    '|',
		EndChar:      '|',
		ProgressChar: '▓',
	}

	bar := &Bar{
		width:          50,
		val:            0,
		maxVal:         maxVal,
		theme:          theme,
		showPercentage: true,
		showTime:       true,
		isFinished:     false,
	}

	// Apply optional arguments
	for _, arg := range kwargs {
		err := arg(bar)
		if err != nil {
			return nil, err
		}
	}

	return bar, nil
}

// update sets the state of the Bar to a new value.
func (b *Bar) update(i int) {

	if b.isFinished {
		return
	}

	// Generate characters to indicate progress
	level := b.width * i / b.maxVal
	progress := strings.Repeat(string(b.theme.ProgressChar), level)
	blanks := strings.Repeat(" ", b.width-level)

	fmt.Printf("\rProgress: %s%s%s%s", string(b.theme.StartChar), progress, blanks, string(b.theme.EndChar))

	if b.showPercentage {
		percentage := 100 * float32(i) / float32(b.maxVal)
		fmt.Printf(" %.2f%%", percentage)
	}

	if b.showTime {
		elapsed := time.Since(b.startTime).Seconds()
		fmt.Printf(" - %.2fs ", elapsed)
	}

	b.val = i
}

// end terminates the progress bar and prints elapsed time.
func (b *Bar) end() {

	if b.isFinished {
		return
	}

	b.update(b.maxVal)
	b.isFinished = true

	elapsed := time.Since(b.startTime)
	fmt.Printf("\nWall time: %f\n", elapsed.Seconds())
}

// Set sets a new value of the Bar
func (b *Bar) Set(i int) {
	b.lock.Lock()
	defer b.lock.Unlock()

	if b.startTime.IsZero() {
		b.Start()
	}

	if i >= b.maxVal {
		b.end()
	} else if i < 0 {
		b.update(0)
	} else {
		b.update(i)
	}
}

// Add adds i to the value of the Bar.
func (b *Bar) Add(i int) {
	b.Set(b.val + i)
}

// Increment adds 1 to the value of the Bar.
func (b *Bar) Increment() {
	b.Add(1)
}

// Start starts the timer on a bar.
// It is implicitly called at the first execution of Set,
// but can be explicitly used earlier if desired.
func (b *Bar) Start() {
	b.startTime = time.Now()
	fmt.Printf("\n")
	b.Set(0)
}

// StartNew creates and starts a new Bar.
func StartNew(maxVal int, kwargs ...func(*Bar) error) (*Bar, error) {
	bar, err := New(maxVal, kwargs...)
	if err != nil {
		return nil, err
	}
	bar.Start()
	return bar, nil
}

// Finish finishes a bar by setting it to max value and terminating it.
func (b *Bar) Finish() {
	b.Set(b.maxVal)
}
