package friend_model

import (
	"time"
)

type Friends struct {
	UserID    int       `gorm:"column:user_id;primary_key"`
	FriendID  int       `gorm:"column:friend_id;NOT NULL"`
	CreatedAt time.Time
}
