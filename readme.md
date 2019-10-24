# A very basic console progress bar for Go

This is a basic exercise to familiarize myself with the Go language.
There are many other more feature-rich progress bars out there, but I think that this is a good practical problem to train on.

## Usage

```Go
package main

import (
	"github.com/lindronics/progressbar"
	"time"
)

func main() {
	b := progressbar.StartNew(100)
	for i := 0; i < 100; i++ {
		b.Increment()
		time.Sleep(20 * time.Millisecond)
    }
    b.Finish()
}
```

### Optional parameters
```Go
// Set display width of the progress bar
b := progressbar.StartNew(100, progressbar.BarWidth(100))

// Set characters for displaying the progress bar
b := progressbar.StartNew(100, progressbar.BarChars('|', '='))

// Determine whether to show a percentage next to the bar
b := progressbar.StartNew(100, progressbar.BarShowPercent(false))

// Determine whether to show elapsed time next to the bar
b := progressbar.StartNew(100, progressbar.BarShowTime(false))
```

## Todo

* [ ] Review thread safety
* [ ] Error handling
* [ ] More extensive theming
* [ ] Testing (maybe CD?)