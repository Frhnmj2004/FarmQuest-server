package server

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Frhnmj2004/FarmQuest-server.git/config"
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/logger"
	"gorm.io/gorm"
)

func Run(ctx context.Context, cfg *config.Config, db *gorm.DB, logger logger.Logger) {
	port := fmt.Sprintf(":%d", cfg.Relayer.Port)

	router := SetupRoutes(ctx, db, logger)

	
	if err := router.Listen(port); err != nil {
		log.Fatalf("Failed to start relayer server: %v", err)
		os.Exit(1)
	}

	log.Printf("listening at 0.0.0.0:%v", port)
}
