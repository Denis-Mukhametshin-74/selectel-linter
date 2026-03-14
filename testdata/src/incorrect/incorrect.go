package main

import "log/slog"

func main() {

	password := "password"
	apiKey := "apiKey"
	token := "token"

	slog.Info("Starting server on port 8080")
	slog.Error("Failed to connect to database")

	slog.Info("запуск сервера")
	slog.Error("ошибка подключения к базе данных")

	slog.Info("server started! �")
	slog.Error("connection failed!!!")
	slog.Warn("warning: something went wrong...")

	slog.Info("user password: " + password)
	slog.Debug("api_key=" + apiKey)
	slog.Info("token: " + token)
}
