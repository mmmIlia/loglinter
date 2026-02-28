package a

import (
	"go.uber.org/zap"
	"log/slog"
)

func main() {
	slog.Info("starting server")
	slog.Error("failed to connect")

	zapLogger := zap.NewExample()
	zapLogger.Debug("all good")

	slog.Info("123 servers started")
	slog.Info("[HTTP] server started")

	slog.Info("Starting server on port 8080")   // want "log message should start with a lowercase letter"
	slog.Error("Failed to connect to database") // want "log message should start with a lowercase letter"

	zapLogger.Warn("Warning: something went wrong") // want "log message should start with a lowercase letter"

	slog.Info("запуск сервера")                    // want "log message must be in English"
	slog.Error("ошибка подключения к базе данных") // want "log message must be in English"

	slog.Info("Падение") // want "log message should start with a lowercase letter" "log message must be in English"

	slog.Info("server started 🚀")                 // want "log message should not contain emojis"
	slog.Error("connection failed!!!")            // want "log message should not contain exclamation or question marks"
	slog.Warn("warning: something went wrong...") // want "log message should not end with punctuation"
	slog.Info("Server запустился! 🚀...")          // want "log message should not contain emojis" "log message should not end with punctuation" "log message should start with a lowercase letter" "log message must be in English" "log message should not contain exclamation or question marks"

	slog.Info("user: created")

	userPassword := "secret123"

	slog.Info("user password saved successfully")

	slog.Info("user password: ")                                     // want "log message should not contain potential sensitive data"
	slog.Info("api_key=")                                            // want "log message should not contain potential sensitive data"
	slog.Info("login attempt: " + userPassword)                      // want "log message should not use variable with potential sensitive data"
	slog.Debug("data", "val", userPassword)                          // want "log message should not use variable with potential sensitive data"
	zapLogger.Info("user created", zap.String("password:", "12345")) // want "log message should not contain potential sensitive data"

	zapLogger.Info("user action", zap.String("UserID", "123"))
}
