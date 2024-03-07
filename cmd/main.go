package main

import (
	"context"
	"log/slog"
	"os"
	"woonbeaj/L0/internal/config"
	"woonbeaj/L0/internal/handlers"
	"woonbeaj/L0/internal/nstreaming"
	memory "woonbeaj/L0/internal/storage"
	"woonbeaj/L0/internal/storage/cache"
	"woonbeaj/L0/internal/storage/postgre"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.MustLoad("config\\config.yaml")
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	db, err := storage.New(cfg.StorageURL)
	if err != nil {
		log.Error("connected db failed: %w", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	cache, err := cache.MustLoad(db)

	srg := memory.NewWithCache(db, cache)

	if err != nil {
		log.Error("cache loading error: %w", err)
	}

	nstreaming.NewConnAndSub(cfg, log, srg)
	
	app := fiber.New()
	app.Static("/", "./ui")
	app.Post("/order", handlers.Order(log, srg))
	err = app.Listen(":8080")
	if err != nil {
		log.Error("server startup failed")
		return
	}

	log.Info("The server has finished its work")
}
