package password

import "time"

type EmailPasswordReset struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id" binding:"required"`
	Phrase     string    `json:"pin" binding:"required"`
	ExpiryDate time.Time `json:"expiry_date" binding:"required"`
}

func (k EmailPasswordReset) IsNoSQLEntity() bool {
	return true
}
