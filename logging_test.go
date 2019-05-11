package logging

import (
	"bytes"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogging(t *testing.T) {
	w := &bytes.Buffer{}
	logrusLogger := logrus.New()
	logrusLogger.SetOutput(w)
	logrusLogger.SetFormatter(&logrus.JSONFormatter{})
	logger := NewLogger(logrusLogger)

	logger.Info("hello")
	logger.WithField("name", "sid").Info("shit")

}

func TestResetLevel(t *testing.T) {
	logrusLogger := logrus.New()

	testcases := map[string]struct {
		level, want logrus.Level
	}{
		"debugLevel": {want: logrus.DebugLevel, level: logrus.TraceLevel},
		"panicLevel": {want: logrus.ErrorLevel, level: logrus.PanicLevel},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			logrusLogger.SetLevel(tc.level)

			resetLevel(logrusLogger)

			if level := logrusLogger.GetLevel(); level != tc.want {
				t.Fatalf("expected level=%v, got=%v", tc.want, level)
			}
		})
	}
}
