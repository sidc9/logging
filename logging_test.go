package logging

import (
	"bufio"
	"bytes"
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sirupsen/logrus"
)

func TestLogging(t *testing.T) {
	w := &bytes.Buffer{}
	logrusLogger := logrus.New()
	logrusLogger.SetOutput(w)
	logrusLogger.SetFormatter(&logrus.JSONFormatter{})
	logger := NewLogger(logrusLogger)

	logger.Info("hello")
	logger.WithField("name", "sid").Info("world")

	wantMsgs := []string{"hello", "world"}
	gotMsgs := make([]string, 0)

	scanner := bufio.NewScanner(w)
	for scanner.Scan() {
		line := scanner.Text()

		var m map[string]string
		if err := json.Unmarshal([]byte(line), &m); err != nil {
			t.Fatal(err)
		}

		if msg, found := m["msg"]; found {
			gotMsgs = append(gotMsgs, msg)
		}
	}

	if diff := cmp.Diff(wantMsgs, gotMsgs); diff != "" {
		t.Fatal(diff)
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
