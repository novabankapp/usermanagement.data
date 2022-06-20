package registration

import (
	"gorm.io/gorm"
	"time"
)

type PhoneVerificationCode struct {
	gorm.Model
	Phone      string    `json:"email" binding:"required"`
	ExpiryDate time.Time `json:"expiry_date" binding:"required"`
	Code       string    `json:"code" binding:"required"`
	Used       bool      `json:"used"`
}

func (e PhoneVerificationCode) IsRDBMSEntity() bool {
	return true
}
