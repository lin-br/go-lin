package main

import (
	"os"
	"time"

	clockface "github.com/lin-br/go-lin/learn-go/16-maths-the-clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
