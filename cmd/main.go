package main

import (
	"edutest/api"
	"edutest/pkg/config"
	logs "edutest/pkg/log"
	"edutest/service"
	"edutest/storage"
	"edutest/storage/postgres"
	"fmt"
)

func main() {
	logger := logs.InitLogger()
	cfg := config.LoadConfig()

	db, err := postgres.Connect(cfg)
	if err != nil {
		logger.Error(fmt.Sprintf("Error is connect postgres: %v", err))
	}


	storage := storage.NewStorage(db, logger, cfg)
	service := service.NewService(storage, logger)
	router := api.Router(service, logger)

	err = router.Run(cfg.EDU_TEST)
	if err != nil {
		logger.Error(fmt.Sprintf("Error is run router: %v", err))
	}
}
