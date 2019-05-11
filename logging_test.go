package logging_test

import (
	"testing"

	"github.com/sidc9/logging"
	"github.com/sirupsen/logrus"
)

func TestLogging(t *testing.T) {
	logrusLogger := logrus.New()
	logger := logging.NewLogger(logrusLogger)

	logger.Info("hello")
	logger.WithField("name", "sid").Info("shit")
}
