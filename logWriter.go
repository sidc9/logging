package logger

import "io"

type logWriter struct {
	*defaultLogger
}

// NewLogWriter returns a wrapped defaultLogger that allows the use of a custom io.Writer.
func NewLogWriter(w io.Writer) *logWriter {
	lw := &logWriter{}
	lw.defaultLogger = NewDefaultLogger()
	lw.log.SetOutput(w)
	return lw
}
