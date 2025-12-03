package main

import (
	"context"
	"fmt"
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
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config/config.yaml"
	}

	cfg := config.LoadConfig(configPath, logger)
	cfg.Storage.SetURI(logger)

	storage := gorm.NewStorage(logger, cfg.Storage.GetURI())
	if err := storage.Ping(context.Background()); err != nil {
		logger.Fatal("failed to ping to storage", zap.Error(err))
	}

	logger.Info(fmt.Sprintf("successfully connected to storage"))

	srv := service.RegisterServices(logger, storage)
	wApp := fiber.New()
	ctrl := controller.NewController(logger, srv, wApp)
	ctrl.ConfigureRoutes()
	logger.Info(fmt.Sprintf("successfully configured routes"))

	quit := make(chan os.Signal, 1)
	go func() {
		if err := wApp.Listen(cfg.Service.Port); err != nil {
			logger.Fatal("Can't shutdown service", zap.Error(err))
			return
		}
	}()

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	if err := storage.Close(context.Background()); err != nil {
		logger.Fatal("failed to close storage", zap.Error(err))
		return
	}

	logger.Info("Database disconnected")
	if err := wApp.Shutdown(); err != nil {
		logger.Info("Failed to stop server")
		return
	}

	logger.Info("Server stopped")
}
