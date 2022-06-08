package registration

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	UserID      string      `json:"user_id" binding:"required"`
	TypeID      int         `json:"type_id"`
	ContactType ContactType `json:"contact_type" gorm:"foreignKey:TypeID"`
	Value       string      `json:"value" binding:"required"`
}
