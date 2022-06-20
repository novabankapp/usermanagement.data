package registration

import (
	"gorm.io/gorm"
)

type UserDetails struct {
	gorm.Model
	UserID        string `json:"user_id" binding:"required"`
	Title         string `json:"firstname" binding:"required"`
	DOB           string `json:"dob" binding:"required"`
	MaritalStatus string `json:"marital_status" binding:"required"`
	Gender        string `json:"gender" binding:"required"`
}

func (e UserDetails) IsRDBMSEntity() bool {
	return true
}
