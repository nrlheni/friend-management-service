package auth_model

import (
	"time"
)

type User struct {
	ID        int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Email     string    `gorm:"column:email;unique;NOT NULL"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time
	Password  string    `gorm:"column:password"`
}