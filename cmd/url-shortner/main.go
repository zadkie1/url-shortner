package main

import (
	"log/slog"
	"os"

	"github.com/zadkie1/url-shortner/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	//config
	cfg := config.MustLoad()

	log := SetupLogger(cfg.Env)

	log.Info("Starting", slog.String("env", cfg.Env))
	log.Info("Debug messages are enabled")

	//init logger

	//init storage

	//init router

	//run server

}
func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	}
	return log

}
