package registration

import (
	"gorm.io/gorm"
)

type UserDetails struct {
	gorm.Model
	UserID        string `json:"user_id" binding:"required" gorm:"type:varchar;not null"`
	Title         string `json:"firstname" binding:"required" gorm:"type:varchar;not null"`
	DOB           string `json:"dob" binding:"required" gorm:"type:varchar;not null"`
	MaritalStatus string `json:"marital_status" binding:"required" gorm:"type:varchar;not null"`
	Gender        string `json:"gender" binding:"required" gorm:"type:varchar;not null"`
}

func (e UserDetails) IsRDBMSEntity() bool {
	return true
}
