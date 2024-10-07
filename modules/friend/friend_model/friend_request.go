package friend_model

import (
	"time"
)

type FriendRequests struct {
	ID          int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	RequesterID int       `gorm:"column:requester_id;NOT NULL"`
	RequesteeID int       `gorm:"column:requestee_id;NOT NULL"`
	Status      string    `gorm:"column:status;default:pending"`
	CreatedAt   time.Time
}
