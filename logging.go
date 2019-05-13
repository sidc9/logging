package logging

import (
	"github.com/sirupsen/logrus"
)

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})

	WithField(key string, value interface{}) Logger
	WithFields(logrus.Fields) Logger
}

type wrappedLogger struct {
	*logrus.Entry
}

func NewLogger(l *logrus.Logger) Logger {
	resetLevel(l)
	return &wrappedLogger{Entry: logrus.NewEntry(l)}
}

func (w *wrappedLogger) Debug(args ...interface{}) {
	w.Entry.Debug(args...)
}
func (w *wrappedLogger) Info(args ...interface{}) {
	w.Entry.Info(args...)
}
func (w *wrappedLogger) Error(args ...interface{}) {
	w.Entry.Error(args...)
}

func (w *wrappedLogger) Debugf(format string, args ...interface{}) {
	w.Entry.Debugf(format, args...)
}
func (w *wrappedLogger) Infof(format string, args ...interface{}) {
	w.Entry.Infof(format, args...)
}
func (w *wrappedLogger) Errorf(format string, args ...interface{}) {
	w.Entry.Errorf(format, args...)
}

func (w *wrappedLogger) WithField(key string, value interface{}) Logger {
	return &wrappedLogger{Entry: w.Entry.WithField(key, value)}
}
func (w *wrappedLogger) WithFields(fields logrus.Fields) Logger {
	return &wrappedLogger{Entry: w.Entry.WithFields(fields)}
}

// resetLevel corrects the logging level. Only Debug, Info and Error levels are allowed.
func resetLevel(l *logrus.Logger) {
	level := l.GetLevel()
	if level < logrus.ErrorLevel {
		l.SetLevel(logrus.ErrorLevel)
	} else if level > logrus.DebugLevel {
		l.SetLevel(logrus.DebugLevel)
	}
}

//func UseRotation(l Logger, filename string) Logger {
//    wl := l.(*wrappedLogger)
//    wl.l.SetOutput(&lumberjack.Logger{
//        Filename: filename,
//        MaxSize:  20,
//    })
//    wl.Entry = logrus.NewEntry(wl.l)
//    return wl
//}
