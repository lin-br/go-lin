package main

import (
	"fmt"
	"io"
	"os"
)

const (
	countdownStart = 3
	finalWord      = "Go!"
)

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprintf(out, finalWord)
}

func main() {
	Countdown(os.Stdout, &DefaultSleeper{})
}
