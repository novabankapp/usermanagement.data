package login

import "time"

type CodeLogin struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id" binding:"required"`
	Code       string    `json:"pin" binding:"required"`
	ExpiryDate time.Time `json:"expiry_date" binding:"required"`
}

func (k CodeLogin) IsNoSQLEntity() bool {
	return true
}
