package apiserver

import (
	"log/slog"
)

func New(opts ...Option) error {
	s := &apiServer{
		logLevel: slog.LevelInfo,
	}

	for _, opt := range opts {
		opt(s)
	}

	//app := fiber.New(fiber.Config{
	//	AppName:      apiPrefix,
	//	IdleTimeout:  ServerIdleTimeout,
	//	ReadTimeout:  ServerReadTimeout,
	//	WriteTimeout: ServerWriteTimeout,
	//})

	return nil
}
