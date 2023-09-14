package dbs

import (
	"fmt"

	"github.com/arsu4ka/todo-auth/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetPostgres(conf *Config) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Name,
		conf.Password,
	)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&models.User{}, &models.Todo{}, &models.Verification{}, &models.Reset{}); err != nil {
		return nil, err
	}

	return db, nil
}
