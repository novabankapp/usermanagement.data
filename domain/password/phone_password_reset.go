package password

import "time"

type PhonePasswordReset struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id" binding:"required"`
	Pin        string    `json:"pin" binding:"required"`
	ExpiryDate time.Time `json:"expiry_date" binding:"required"`
}
