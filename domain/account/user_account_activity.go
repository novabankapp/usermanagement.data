package account

import (
	"github.com/gocql/gocql"
	"time"
)

type UserAccountActivity struct {
	ID           gocql.UUID `json:"id"`
	AccountID    string     `json:"account_id"`
	Activity     string     `json:"activity"`
	IpAddress    string     `json:"ip_address"`
	ActivityDate time.Time  `json:"created_at"`
}

func (k UserAccountActivity) IsNoSQLEntity() bool {
	return true
}
