package registration

import "gorm.io/gorm"

type UserEmployment struct {
	gorm.Model
	UserID         string `json:"user_id" binding:"required" gorm:"type:varchar;not null"`
	NameOfEmployer string `json:"name_of_employer" binding:"required" gorm:"type:varchar;not null"`
	Industry       string `json:"industry" binding:"required" gorm:"type:varchar;not null"`
}

func (e UserEmployment) IsRDBMSEntity() bool {
	return true
}
