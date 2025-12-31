package db

import "io"

// Create a Tape type to encapsulate the
// Write method, giving me control to
// write from the beginning
type tape struct {
	file io.ReadWriteSeeker
}

func (t *tape) Write(p []byte) (n int, err error) {
	_, _ = t.file.Seek(0, io.SeekStart)
	return t.file.Write(p)
}
