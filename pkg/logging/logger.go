package logging

import (
	"log/slog"
	"os"
)

type Mode string
type Logger = *slog.Logger

const (
	Local Mode = "local"
	Dev   Mode = "dev"
	Prod  Mode = "prod"
)

func getLocalLogger() Logger {
	return slog.New(NewPrettyHandler(&slog.HandlerOptions{Level: slog.LevelDebug}))
}

func getDevLogger() Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
}

func getProdLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
}

func GetLogger(mode Mode) Logger {
	var log Logger

	switch mode {
	case Local:
		log = getLocalLogger()
	case Dev:
		log = getDevLogger()
	case Prod:
		log = getProdLogger()
	}

	return log
}
