package service

import (
	"edutest/pkg/config"
	"edutest/storage"
	"log/slog"
)

type Service struct {
	Storage storage.Storage
	Log     *slog.Logger
	Cfg config.Config
}

func NewService(storage storage.Storage, log *slog.Logger, cfg config.Config) Service {
	return Service{
		Storage: storage,
		Log:     log,
		Cfg: cfg,
	}
}
