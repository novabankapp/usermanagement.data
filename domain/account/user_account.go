package account

import (
	"github.com/gocql/gocql"
	"time"
)

type UserAccount struct {
	ID        gocql.UUID `json:"id"`
	UserID    string     `json:"user_id" binding:"required"`
	CreatedAt time.Time  `json:"created_at"`
	IsLocked  bool       `json:"is_locked"`
	IsActive  bool       `json:"is_active"`
}

func (k UserAccount) IsNoSQLEntity() bool {
	return true
}
