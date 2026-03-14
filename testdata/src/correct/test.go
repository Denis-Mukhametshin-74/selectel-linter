package main

import "log/slog"

func main() {
	slog.Info("starting server on port 8080")
	slog.Error("failed to connect to database")

	slog.Info("starting server")
	slog.Error("failed to connect to database")

	slog.Info("server started")
	slog.Error("connection failed")
	slog.Warn("something went wrong")

	slog.Info("user authenticated successfully")
	slog.Debug("api request completed")
	slog.Info("token validated")
}
