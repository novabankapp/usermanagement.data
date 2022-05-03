package domain

import "github.com/shopspring/decimal"

type UserIncome struct {
	UserID        string          `json:"user_id" binding:"required"`
	Source        string          `json:"source" binding:"required"`
	MonthlyIncome decimal.Decimal `json:"monthly_income" sql:"type:decimal(20,2);"`
	ProofOfSource string          `json:"proof_of_source" binding:"required"`
}
