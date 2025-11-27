package main

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"okami-qstn-bnk/internal/config"
	controller "okami-qstn-bnk/internal/controller/http/fiber"
	"okami-qstn-bnk/internal/service"
	"okami-qstn-bnk/internal/storage/gorm"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger, _ := zap.NewDevelopment()

	cfg := config.LoadConfig("config/config.yaml", logger)
	cfg.Storage.SetURI(logger)

	storage := gorm.NewStorage(logger, cfg.Storage.GetURI())

	qsrv, tsrv := service.RegisterServices(logger, storage)
	wApp := fiber.New()
	ctrl := controller.NewController(logger, qsrv, tsrv, wApp)
	ctrl.ConfigureRoutes()

	quit := make(chan os.Signal, 1)

	go func() {
		if err := wApp.Listen(cfg.Service.Port); err != nil {
			logger.Fatal("Can't shutdown service", zap.Error(err))
		}
	}()

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	if err := wApp.Shutdown(); err != nil {
		logger.Info("Failed to stop server")
		return
	}

	logger.Info("Server stopped")
}
