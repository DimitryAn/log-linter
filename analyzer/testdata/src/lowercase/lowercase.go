package lowercase

import (
	"log"
	"log/slog"
)

func Lowercase() {
	log.Fatal("server 88 98 8080")

	log.Fatal("Starting server on port 8080") // want "message must start with lowercase letter"

	slog.Error("Failed to connect to database") // want "message must start with lowercase letter"

	log.Fatal("SERVER" + "s" + "S") // want "message must start with lowercase letter" "message must start with lowercase letter"

	slog.Info("failed to coNnect to database")

	log.Fatal("sasa222", 22) // want "message must contains only english letters, <space> and digits" "message must contains only english letters, <space> and digits"
}
