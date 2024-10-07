package friend_model

import (
	"friends-management-api/modules/auth/auth_model"
	"time"
)

type FriendRequests struct {
	ID          int       			`gorm:"column:id;primary_key;AUTO_INCREMENT"`
	RequesterID int       			`gorm:"column:requester_id;NOT NULL"`
	RequesteeID int       			`gorm:"column:requestee_id;NOT NULL"`
	Status      string    			`gorm:"column:status;default:pending"`
	CreatedAt   time.Time
	Requester   auth_model.User     `gorm:"foreignKey:RequesterID;references:ID"`
	Requestee   auth_model.User     `gorm:"foreignKey:RequesteeID;references:ID"`
}
