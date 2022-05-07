package registration

import (
	"gorm.io/gorm"
	"time"
)

type UserOneTimePin struct {
	gorm.Model
	UserID     string    `json:"user_id" binding:"required"`
	Pin        string    `json:"pin" binding:"required"`
	ExpiryDate time.Time `json:"expiry_date" binding:"required"`
}
