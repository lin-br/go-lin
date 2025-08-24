package main

import (
	"os"
	"time"

	"github.com/lin-br/go-lin/learn-go/16-maths-the-clockface/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
