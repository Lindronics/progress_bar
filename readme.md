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
    b := progressbar.StartNew(50)
    for i := 0; i < 50; i++ {
        b.Increment()
        time.Sleep(20 * time.Millisecond)
    }
    b.Finish()
}
```

### Example output
```
Progress: |▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓| 100.00% - 0.63s 
Wall time: 0.634013
```

### Optional parameters
```Go
// Set display width of the progress bar
b := progressbar.StartNew(100, progressbar.BarWidth(100))

// Determine whether to show a percentage next to the bar
b := progressbar.StartNew(100, progressbar.BarShowPercent(false))

// Determine whether to show elapsed time next to the bar
b := progressbar.StartNew(100, progressbar.BarShowTime(false))
```

### Theming
```Go
theme := Theme{
    StartChar:    '[',
    EndChar:      ']',
    ProgressChar: '=',
}

b := progressbar.StartNew(100, progressbar.BarTheme(theme)

```

## Todo

* [ ] Review thread safety
* [ ] Error handling
* [x] More extensive theming
* [ ] Testing (maybe CD?)