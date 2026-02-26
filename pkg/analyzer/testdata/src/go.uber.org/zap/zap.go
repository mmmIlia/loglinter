package zap

// Создаем пустышку Logger
type Logger struct{}

// NewExample возвращает наш фейковый логгер (мы использовали его в main.go)
func NewExample() *Logger {
	return &Logger{}
}

// Прописываем сигнатуры методов. Нам важно только то, что первый аргумент - строка (msg).
// Остальное линтеру не важно.
func (l *Logger) Info(msg string, fields ...any) {}
func (l *Logger) Error(msg string, fields ...any) {}
func (l *Logger) Warn(msg string, fields ...any) {}
func (l *Logger) Debug(msg string, fields ...any) {}
func (l *Logger) Fatal(msg string, fields ...any) {}