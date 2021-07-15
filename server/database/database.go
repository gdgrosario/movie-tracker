package database

import (
	"github.com/gdgrosario/movieland/database/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Connect : Database connect
func Connect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./database/data.db")
	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	models.Migrate(db)

	return db
}
