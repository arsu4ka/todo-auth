package dbs

import (
	"fmt"

	"github.com/arsu4ka/todo-auth/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostgres(conf *Config) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Name,
		conf.Password,
	)

	db, err := gorm.Open(postgres.Open(connectionString))
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&models.User{}, &models.Todo{}); err != nil {
		return nil, err
	}

	return db, nil
}
