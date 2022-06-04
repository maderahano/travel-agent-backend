package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"travel-agent-backend/config"
	"travel-agent-backend/internal/model"
)

func InitDB(conf config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
	)

	DB, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err.Error())
	}

	initMigrate(DB)
	return DB
}

func initMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}
