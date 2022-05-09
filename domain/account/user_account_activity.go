package account

import "time"

type UserAccountActivity struct {
	ID           string    `json:"id"`
	AccountID    string    `json:"account_id"`
	Activity     string    `json:"activity"`
	ActivityDate time.Time `json:"created_at"`
}
