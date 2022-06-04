package registration

import (
	"gorm.io/gorm"
	"time"
)

type EmailVerificationCode struct {
	gorm.Model
	Email      string    `json:"email" binding:"required"`
	ExpiryDate time.Time `json:"expiry_date" binding:"required"`
	Code       string    `json:"code" binding:"required"`
}