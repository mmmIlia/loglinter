package a

import (
	"log/slog"
	"go.uber.org/zap"
)

func main() {
	slog.Info("starting server")
	slog.Error("failed to connect")
	
	zapLogger := zap.NewExample()
	zapLogger.Debug("all good")
	
	slog.Info("123 servers started")
	slog.Info("[HTTP] server started")


	slog.Info("Starting server on port 8080") // want "log message should start with a lowercase letter"
	slog.Error("Failed to connect to database") // want "log message should start with a lowercase letter"
	
	zapLogger.Warn("Warning: something went wrong") // want "log message should start with a lowercase letter"


	slog.Info("запуск сервера") // want "log message must be in English"
	slog.Error("ошибка подключения к базе данных") // want "log message must be in English"
	
	slog.Info("Падение") // want "log message should start with a lowercase letter" "log message must be in English"


	slog.Info("server started 🚀") // want "log message should not contain emojis"
	slog.Error("connection failed!!!") // want "log message should not contain exclamation or question marks"
	slog.Warn("warning: something went wrong...") // want "log message should not end with punctuation"
	
	slog.Info("user: created") 
}