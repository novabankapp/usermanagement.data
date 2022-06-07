package account

import "time"

type UserAccount struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	IsLocked  bool      `json:"is_locked"`
	IsActive  bool      `json:"is_active"`
}
