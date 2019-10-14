package logger

import "os"

type fileLogger struct {
	*logWriter

	age  int
	size int
}

func NewFileLogger(filename string, opts ...FileOptions) (*fileLogger, error) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	fl := &fileLogger{
		logWriter: NewLogWriter(f),
	}

	for _, opt := range opts {
		opt(fl)
	}

	return fl
}

type FileOptions func(*fileLogger)

func MaxAge(age int) FileOptions {
	return func(fl *fileLogger) {
		fl.age = age
	}
}

func MaxSize(size int) FileOptions {
	return func(fl *fileLogger) {
		fl.size = size
	}
}
