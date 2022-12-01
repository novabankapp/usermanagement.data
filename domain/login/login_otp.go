package login

import (
	"github.com/gocql/gocql"
	"time"
)

type OtpLogin struct {
	ID         gocql.UUID `json:"id"`
	UserID     string     `json:"user_id" binding:"required"`
	Pin        string     `json:"pin" binding:"required"`
	ExpiryDate time.Time  `json:"expiry_date" binding:"required"`
}

func (k OtpLogin) IsNoSQLEntity() bool {
	return true
}
