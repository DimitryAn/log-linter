package sensetive

import (
	"log"
	"log/slog"
	"os"
)

func Lowercase() {

	var password, apiKey, token, myapp string

	slog.Error("user password: " + password) // want "message must contains only english letters, <space> and digits" "potential sensitive data, please rename variable"
	slog.Info("api_key=" + apiKey)           // want "message must contains only english letters, <space> and digits" "potential sensitive data, please rename variable"
	slog.Info("token: " + token)             // want "message must contains only english letters, <space> and digits" "potential sensitive data, please rename variable"
	log.Panic("username" + password)         // want "potential sensitive data, please rename variable"

	log.Fatal(password) // want "potential sensitive data, please rename variable"
	log.Fatal(myapp)

	os.Create("password")
	os.Getenv("pass")

	log.Fatal("Server started" + myapp) // want "message must start with lowercase letter"
}
