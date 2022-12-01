package registration

import (
	"gorm.io/gorm"
	"time"
)

type UserOneTimePin struct {
	gorm.Model
	UserID     string    `json:"user_id" binding:"required" gorm:"type:varchar;not null"`
	Pin        string    `json:"pin" binding:"required" gorm:"type:varchar;not null"`
	ExpiryDate time.Time `json:"expiry_date" binding:"required"`
}

func (e UserOneTimePin) IsRDBMSEntity() bool {
	return true
}
