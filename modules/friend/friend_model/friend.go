package friend_model

import (
	"friends-management-api/modules/auth/auth_model"
	"time"
)

type Friends struct {
	UserID    int       		`gorm:"column:user_id;primary_key"`
	FriendID  int       		`gorm:"column:friend_id;NOT NULL"`
	CreatedAt time.Time
	User      auth_model.User    `gorm:"foreignKey:UserID;references:ID"`
	Friend    auth_model.User    `gorm:"foreignKey:FriendID;references:ID"`
}
