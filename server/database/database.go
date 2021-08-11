package database

import (
	"fmt"
	"log"

	"github.com/gdgrosario/movieland/database/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// ConnectDB Make the connection to the database and return the same connection
func ConnectDB() *gorm.DB{
  const (
    host = "localhost"
    user = "golang"
    port = 5432
    password = "golang"
    name = "moviland"
  )

  db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name))

  if err != nil {
    log.Fatalf("Error in connect the DB %v", err)
    return nil
  }

  models.Migrate(db)

  log.Println("DB connect")

  return db
}

