package registration

import (
	"gorm.io/gorm"
	"time"
)

type EmailVerificationCode struct {
	gorm.Model
	Email      string    `json:"email" binding:"required" gorm:"type:varchar;not null"`
	ExpiryDate time.Time `json:"expiry_date" binding:"required"`
	Used       bool      `json:"used"`
	Code       string    `json:"code" binding:"required" gorm:"type:varchar;not null"`
}

func (e EmailVerificationCode) IsRDBMSEntity() bool {
	return true
}
