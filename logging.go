package logging

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})

	WithField(key string, value interface{}) Logger
}

type wrappedLogger struct {
	*logrus.Entry
	//l *logrus.Logger
}

// checkLevel corrects the logging level. Only Debug, Info and Error levels are allowed.
func checkLevel(l *logrus.Logger) {
	level := l.GetLevel()
	if level < logrus.ErrorLevel {
		l.SetLevel(logrus.ErrorLevel)
	} else if level > logrus.DebugLevel {
		l.SetLevel(logrus.DebugLevel)
	}
	fmt.Println(level, l.GetLevel())
}

func NewLogger(l *logrus.Logger) Logger {
	checkLevel(l)
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

func (w *wrappedLogger) WithField(key string, value interface{}) Logger {
	return &wrappedLogger{Entry: w.Entry.WithField(key, value)}
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
