package logger

import (
	"github.com/sirupsen/logrus"
)

type defaultLogger struct {
	log    *logrus.Logger
	fields map[string]interface{}
}

// NewDefaultLogger returns a defaultLogger. This is the most basic logger. It writes
// to os.Stdout by default, and the default level is set to Info. It uses a logrus.Log
// internally to do the actual logging.
func NewDefaultLogger() *defaultLogger {
	dl := &defaultLogger{}
	dl.log = logrus.New()
	dl.log.SetFormatter(&logrus.TextFormatter{})
	dl.log.SetLevel(logrus.InfoLevel)
	return dl
}

// SetLevel lets you set the level of the Logger. The allowed levels are
// Debug, Info and Error.
func (dl *defaultLogger) SetLevel(level Level) {
	var lvl logrus.Level
	switch level {
	case Debug:
		lvl = logrus.DebugLevel
	case Info:
		lvl = logrus.InfoLevel
	case Error:
		lvl = logrus.ErrorLevel
	}
	dl.log.SetLevel(logrus.Level(lvl))
}

// Debug writes a log at the debug level.
func (dl *defaultLogger) Debug(msg string) {
	dl.log.WithFields(logrus.Fields(dl.fields)).Debug(msg)
}

// Info writes a log at the info level.
func (dl *defaultLogger) Info(msg string) {
	dl.log.WithFields(logrus.Fields(dl.fields)).Info(msg)
}

// Error writes a log at the error level.
func (dl *defaultLogger) Error(msg string) {
	dl.log.WithFields(logrus.Fields(dl.fields)).Error(msg)
}

// WithFields returns a new logger with the provided fields added to the existing fields.
// If there are fields with the same names, the old fields are overwritten.
func (dl *defaultLogger) WithFields(fields map[string]interface{}) Logger {
	ndl := &defaultLogger{
		log:    dl.log,
		fields: make(map[string]interface{}),
	}
	for k, v := range fields {
		ndl.fields[k] = v
	}
	for k, v := range dl.fields {
		ndl.fields[k] = v
	}
	return ndl
}
