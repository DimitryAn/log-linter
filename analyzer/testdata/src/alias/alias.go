package alias

import (
	mylg "log"
	slg "log/slog"
	sss "os"
)

func Alias() {

	var password, apiKey, token, myapp string

	slg.Error("user password: " + password) // want "message must contains only english letters, <space> and digits" "potential sensitive data, please rename variable"
	slg.Info("api_key=" + apiKey)           // want "message must contains only english letters, <space> and digits" "potential sensitive data, please rename variable"
	slg.Info("token: " + token)             // want "message must contains only english letters, <space> and digits" "potential sensitive data, please rename variable"
	mylg.Panic("username" + password)       // want "potential sensitive data, please rename variable"

	mylg.Fatal(password) // want "potential sensitive data, please rename variable"
	mylg.Fatal(myapp)

	sss.Create("password")
	sss.Getenv("pass")

	mylg.Fatal("Server started" + myapp) // want "message must start with lowercase letter"

	mylg.Fatal("запуск сервера")                  // want "message must contains only english letters, <space> and digits"
	slg.Error("ошибка подключения к базе данных") // want "message must contains only english letters, <space> and digits"
	mylg.Fatal(" сава")                           // want "message must start with lowercase letter" "message must contains only english letters, <space> and digits"
}
