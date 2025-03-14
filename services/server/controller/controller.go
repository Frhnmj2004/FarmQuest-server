package controller

import (
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/logger"
	"gorm.io/gorm"
)

type Controller struct {
	db     *gorm.DB
	logger logger.Logger
}

func NewBaseController(db *gorm.DB, logger logger.Logger) *Controller {
	return &Controller{db: db, logger: logger}
}
