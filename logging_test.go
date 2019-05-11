package logging

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
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

	msgs := make([]string, 0)
	scanner := bufio.NewScanner(w)
	for scanner.Scan() {
		line := scanner.Text()

		var m map[string]interface{}
		if err := json.Unmarshal(line, &m); err != nil {
			t.Fatal(err)
		}

		if msg, found := m["msg"]; found {
			msgs = append(msgs, msg)
		}
	}

	for _, msg := range msgs {
		fmt.Println(msg)
	}
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
