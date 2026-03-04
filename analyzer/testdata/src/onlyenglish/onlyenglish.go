package onlyenglish

import (
	"log"
	"log/slog"
)

func Lowercase() {
	log.Fatal("запуск сервера") // want "message must contains only english letters, <space> and digits"

	slog.Error("ошибка подключения к базе данных") // want "message must contains only english letters, <space> and digits"
}
