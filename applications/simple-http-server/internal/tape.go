package poker

import (
	"io"
	"os"
)

// Create a Tape type to encapsulate the
// Write method, giving me control to
// write from the beginning
type tape struct {
	// Change to os.File and truncate the file before writing
	file *os.File
}

func (t *tape) Write(p []byte) (n int, err error) {
	// Change to os.File and truncate the file before writing
	_ = t.file.Truncate(0)
	_, _ = t.file.Seek(0, io.SeekStart)
	return t.file.Write(p)
}
