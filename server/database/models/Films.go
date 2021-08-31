package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Post : Post Model
type Film struct {
	gorm.Model
	ID uint64 `gorm:"primary_key"` // auto increment uint default true

	Content  string
  Category string
	Date     time.Time
	Title    string
  Views    int
}
