package login

import "time"

type UserLoginAttempt struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	AttemptDate time.Time `json:"attempt_date"`
	IPAddress   string    `json:"ip_address"`
	Attempts    int       `json:"attempts"`
}

func (k UserLoginAttempt) IsNoSQLEntity() bool {
	return true
}
