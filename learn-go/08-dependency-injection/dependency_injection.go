package dependency_injection

import (
	"fmt"
	"io"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	/* If this file were in the main package, it would work properly with the Go run command. */
	Greet(os.Stdout, "Lin")
}
