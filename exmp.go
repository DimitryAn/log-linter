package basicTest

import (
	"log"
	"log/slog"
)

func Example() {
	log.Fatal("sasa222", 22) // want "message must start with lowercase letter"
	slog.Info("Ads")
}
