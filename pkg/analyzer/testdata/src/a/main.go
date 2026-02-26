package a

import (
	"log/slog"
	"go.uber.org/zap"
)

type MyCustomLogger struct{}

func (l *MyCustomLogger) Info(msg string) {}

func main() {
	slog.Info("this is a test") // want "found log message in log/slog.Info"

	logger := slog.Default()
	logger.Error("another test") // want "found log message in log/slog.Error"

	zapLogger := zap.NewExample()
	zapLogger.Warn("zap test") // want "found log message in go.uber.org/zap.Warn"

	fakeLogger := &MyCustomLogger{}
	fakeLogger.Info("Should be ignored")
}