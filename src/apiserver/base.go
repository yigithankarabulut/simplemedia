package apiserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/apiserver/routes"

	"log/slog"
	"time"
)

const (
	ContextCancelTimeout = 5 * time.Second
	ShutdownTimeout      = 2 * time.Second
	Port                 = ":8080"
)

type apiServer struct {
	app      *fiber.App
	handlers []routes.Handler
	logLevel slog.Level
	logger   *slog.Logger
	port     string
}

type Option func(*apiServer)

func WithLogger(logger *slog.Logger) Option {
	return func(s *apiServer) {
		s.logger = logger
	}
}

func WithLogLevel(level string) Option {
	return func(s *apiServer) {
		var logLevel slog.Level

		switch level {
		case "DEBUG":
			logLevel = slog.LevelDebug
		case "INFO":
			logLevel = slog.LevelInfo
		case "WARN":
			logLevel = slog.LevelWarn
		case "ERROR":
			logLevel = slog.LevelError
		default:
			logLevel = slog.LevelInfo
		}
		s.logLevel = logLevel
	}
}
