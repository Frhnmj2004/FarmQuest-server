package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Frhnmj2004/FarmQuest-server.git/config"
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/db"
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/logger"
	"github.com/Frhnmj2004/FarmQuest-server.git/services/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	ctx := context.Background()
	cfg := config.LoadConfig()
	logger := logger.NewLogger(cfg.Environment)
	db := db.NewDB(cfg)

	// Listen for OS signals (like Ctrl+C)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	server.Run(ctx, cfg, db, logger)

	// Block until we receive a signal
	sigReceived := <-signalChan
	log.Printf("Received signal: %v", sigReceived)
}
