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

	db.AutoMigrate(&models.User{})
	return db, nil
}

func GetPostgresNoAuth(conf *Config) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s dbname=%s sslmode=disable",
		conf.Host,
		conf.Name,
	)

	db, err := gorm.Open(postgres.Open(connectionString))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{}, &models.ToDo{})
	return db, nil
}
