package xco

import (
	"io"
	"log"
)

type writeLogger struct {
	log *log.Logger
	w   io.Writer
}

func (l *writeLogger) Write(p []byte) (n int, err error) {
	n, err = l.w.Write(p)
	if err == nil {
		l.log.Printf("|> %s", p[0:n])
	} else {
		l.log.Printf("|> %s: %v", p[0:n], err)
	}
	return
}

// newWriteLogger returns a writer that behaves like w except that it
// logs the string written.
func newWriteLogger(log *log.Logger, w io.Writer) io.Writer {
	return &writeLogger{log, w}
}

type readLogger struct {
	log *log.Logger
	r   io.Reader
}

func (l *readLogger) Read(p []byte) (n int, err error) {
	n, err = l.r.Read(p)
	if err == nil {
		l.log.Printf("<| %s", p[0:n])
	} else {
		l.log.Printf("<| %s: %v", p[0:n], err)
	}
	return
}

// newReadLogger returns a reader that behaves like r except that it
// logs the string read.
func newReadLogger(log *log.Logger, r io.Reader) io.Reader {
	return &readLogger{log, r}
}
