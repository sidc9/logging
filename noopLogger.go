package logger

type noopLogger struct {
}

// NewNoopLogger returns an instance of noopLogger (a no-op logger).
// It replaces the methods of Logger interface with methods that do nothing.
// It can be used in tests where logs are not required.
func NewNoopLogger() *noopLogger {
	return &noopLogger{}
}

func (*noopLogger) Debug(msg string) {}

func (*noopLogger) Info(msg string) {}

func (*noopLogger) Error(msg string) {}

func (n *noopLogger) WithFields(fields map[string]interface{}) Logger {
	return n
}
