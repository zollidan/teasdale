package database

import (
	"github.com/zollidan/teasdale/config"
	"github.com/zollidan/teasdale/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.DBURL), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// AutoMigrate всех моделей
	if err := autoMigrate(db); err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	return db
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Artist{},
		&models.Genre{},
		&models.Album{},
		&models.Track{},
		&models.TrackGenre{},
		&models.Like{},
		&models.Comment{},
		&models.Review{},
		&models.News{},
		&models.TrendingScore{},
	)
}