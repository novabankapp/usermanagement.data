package registration

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type UserIncome struct {
	gorm.Model
	UserID        string          `json:"user_id" binding:"required"`
	Source        string          `json:"source" binding:"required"`
	MonthlyIncome decimal.Decimal `json:"monthly_income" sql:"type:decimal(20,2);"`
	ProofOfSource string          `json:"proof_of_source" binding:"required"`
}
