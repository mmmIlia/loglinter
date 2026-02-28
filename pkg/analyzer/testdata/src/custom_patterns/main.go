package custom_patterns

import (
	"log/slog"
)

func main() {
	userEmail := "test@test.com"

	slog.Info("user email: ")               // want "log message should not contain potential sensitive data"
	slog.Info("failed login: " + userEmail) // want "log message should not use variable with potential sensitive data"
}
