package registration

import (
	"gorm.io/gorm"
	"time"
)

type PhoneVerificationCode struct {
	gorm.Model
	Phone      string    `json:"email" binding:"required" gorm:"type:varchar;not null"`
	ExpiryDate time.Time `json:"expiry_date" binding:"required"`
	Code       string    `json:"code" binding:"required" gorm:"type:varchar;not null"`
	Used       bool      `json:"used"`
}

func (e PhoneVerificationCode) IsRDBMSEntity() bool {
	return true
}
