package repositories

import (
	"Telegram-Store/database"

	"gorm.io/gorm"
)

var DB *gorm.DB = database.Init()
