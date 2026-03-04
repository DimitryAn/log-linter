package nonenglish

import (
	"log"
	"log/slog"
)

func Nonenglish() {

	log.Fatal("lk дд") // want "message must contains only english letters, <space> and digits"

	log.Fatal("Low рус") // want "message must contains only english letters, <space> and digits" "message must start with lowercase letter"

	slog.Info("!!!") // want "message must contains only english letters, <space> and digits" "message must start with lowercase letter"

	slog.Info(".") // want "message must contains only english letters, <space> and digits" "message must start with lowercase letter"
}
