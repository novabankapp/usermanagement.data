package registration

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	UserID      string      `json:"user_id" binding:"required"`
	ContactType ContactType `json:"contact_type"`
	Value       string      `json:"value" binding:"required"`
}
