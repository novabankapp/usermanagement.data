package registration

import "gorm.io/gorm"

type UserEmployment struct {
	gorm.Model
	UserID         string `json:"user_id" binding:"required"`
	NameOfEmployer string `json:"name_of_employer" binding:"required"`
	Industry       string `json:"industry" binding:"required"`
}
