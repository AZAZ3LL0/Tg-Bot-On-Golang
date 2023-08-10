package repositories

import (
	"Telegram-Store/storage/database"

	"gorm.io/gorm"
)

var DB *gorm.DB = database.Init()
