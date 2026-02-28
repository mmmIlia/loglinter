package zap

type Logger struct{}

func NewExample() *Logger {
	return &Logger{}
}

func (l *Logger) Info(msg string, fields ...any)  {}
func (l *Logger) Error(msg string, fields ...any) {}
func (l *Logger) Warn(msg string, fields ...any)  {}
func (l *Logger) Debug(msg string, fields ...any) {}
func (l *Logger) Fatal(msg string, fields ...any) {}

func String(key string, val string) any { return nil }
