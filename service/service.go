package service

import (
	"edutest/storage"
	"log/slog"
)

type Service struct {
	Storage storage.Storage
	Log     *slog.Logger
}

func NewService(storage storage.Storage, log *slog.Logger) Service {
	return Service{
		Storage: storage,
		Log:     log,
	}
}
