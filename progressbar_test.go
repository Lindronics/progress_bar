package progressbar

import (
	"os"
	"testing"
)

func TestBarWidth(t *testing.T) {
	cases := []struct {
		value     int
		throwsErr bool
	}{
		{10, false},
		{100, false},
		{0, true},
		{-10, true},
	}

	for _, c := range cases {
		bar, err := New(10)
		err = BarWidth(c.value)(bar)
		if (err != nil) != c.throwsErr {
			t.Errorf("Input: %d. Error expected: %t. Error thrown: %t", c.value, c.throwsErr, (err != nil))
		}
	}
}

func TestBarTheme(t *testing.T) {
	cases := []struct {
		value     Theme
		throwsErr bool
	}{
		{Theme{'|', '|', '▓'}, false},
		{Theme{}, true},
		{Theme{'|', 0, 0}, true},
		{Theme{0, '|', 0}, true},
		{Theme{0, 0, '▓'}, true},
	}

	for _, c := range cases {
		bar, err := New(10)
		err = BarTheme(c.value)(bar)
		if (err != nil) != c.throwsErr {
			t.Errorf("Input: %d. Error expected: %t. Error thrown: %t", c.value, c.throwsErr, (err != nil))
		}
	}
}

func TestNew(t *testing.T) {

	// Case 1
	_, err := New(-10)
	if err == nil {
		t.Errorf("Error expected! Invalid maxVal %d.", -10)
	}

	// Case 2
	_, err = New(-10, BarWidth(-20))
	if err == nil {
		t.Errorf("Error expected! Invalid BarWidth %d.", -20)
	}

	// Case 3
	theme := Theme{}
	_, err = New(-10, BarTheme(theme))
	if err == nil {
		t.Errorf("Error expected! Invalid BarTheme %d.", theme)
	}

	// Case 4
	_, err = New(10, BarWidth(50))
	if err != nil {
		t.Errorf("Error thrown, no error expected!")
	}
}

func TestStartNew(t *testing.T) {

	os.Stdout = nil

	// Negative maxVal
	_, err := StartNew(-10)
	if err == nil {
		t.Errorf("Error expected! Invalid maxVal %d.", -10)
	}

	// Error in varArgs
	_, err = StartNew(-10, BarWidth(-20))
	if err == nil {
		t.Errorf("Error expected! Invalid BarWidth %d.", -20)
	}

	// Error in varArgs
	theme := Theme{}
	_, err = StartNew(-10, BarTheme(theme))
	if err == nil {
		t.Errorf("Error expected! Invalid BarTheme %d.", theme)
	}

	// Positive case
	_, err = StartNew(10, BarWidth(50))
	if err != nil {
		t.Errorf("Error thrown, no error expected!")
	}
}

func TestSet(t *testing.T) {

	os.Stdout = nil

	// Positive case
	bar, _ := StartNew(10)
	bar.Set(5)
	if bar.isFinished {
		t.Errorf("Bar should not be finished!")
	}
	if bar.val != 5 {
		t.Errorf("Invalid value! Is: %d, should be: %d", bar.val, 5)
	}

	// Too large values should end the bar
	bar, _ = StartNew(10)
	bar.Set(20)
	if !bar.isFinished {
		t.Errorf("Bar should be finished!")
	}

	// Negative values should be interpreted as 0
	bar, _ = StartNew(10)
	bar.Set(-10)
	if bar.isFinished {
		t.Errorf("Bar should not be finished!")
	}
	if bar.val != 0 {
		t.Errorf("Invalid value! Is: %d, should be: %d", bar.val, 0)
	}
}

func TestIncrement(t *testing.T) {

	os.Stdout = nil

	// Positive case
	bar, _ := StartNew(10)
	bar.Increment()
	if bar.isFinished {
		t.Errorf("Bar should not be finished!")
	}
	if bar.val != 1 {
		t.Errorf("Invalid value! Is: %d, should be: %d", bar.val, 1)
	}

	// Too large values should end the bar
	bar, _ = StartNew(10)
	bar.Set(9)
	bar.Increment()
	if !bar.isFinished {
		t.Errorf("Bar should be finished!")
	}
}

func TestAdd(t *testing.T) {

	os.Stdout = nil

	// Positive case
	bar, _ := StartNew(10)
	bar.Add(5)
	if bar.isFinished {
		t.Errorf("Bar should not be finished!")
	}
	if bar.val != 5 {
		t.Errorf("Invalid value! Is: %d, should be: %d", bar.val, 5)
	}

	// Too large values should end the bar
	bar, _ = StartNew(10)
	bar.Set(9)
	bar.Add(5)
	if !bar.isFinished {
		t.Errorf("Bar should be finished!")
	}

	// Negative values should be interpreted as 0
	bar, _ = StartNew(10)
	bar.Set(2)
	bar.Add(-5)
	if bar.isFinished {
		t.Errorf("Bar should not be finished!")
	}
	if bar.val != 0 {
		t.Errorf("Invalid value! Is: %d, should be: %d", bar.val, 0)
	}
}
