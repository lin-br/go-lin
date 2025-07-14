package main

import (
	"fmt"
	"io"
	"iter"
	"os"
)

const (
	countdownStart = 3
	finalWord      = "Go!"
)

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := range countDownFrom(countdownStart) {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func countDownFrom(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

func main() {
	Countdown(os.Stdout, &DefaultSleeper{})
}
