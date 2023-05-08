package dbs

import (
	"github.com/arsu4ka/todo-auth/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetSqlite(DBPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(DBPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{}, &models.ToDo{})
	return db, nil
}
