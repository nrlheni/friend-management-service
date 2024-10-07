package config

import (
	"fmt"
	"friends-management-api/modules/auth/auth_model"
	"friends-management-api/modules/friend/friend_model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(cfg *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.DbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&auth_model.User{})
	db.AutoMigrate(&friend_model.Friends{})
	db.AutoMigrate(&friend_model.Blocks{})
	db.AutoMigrate(&friend_model.FriendRequests{})

	return db, nil
}
