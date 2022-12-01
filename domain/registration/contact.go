package registration

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	UserID      string      `json:"user_id" binding:"required" gorm:"type:varchar;not null"`
	TypeID      int         `json:"type_id" gorm:"type:int;not null"`
	ContactType ContactType `json:"contact_type" gorm:"foreignKey:TypeID"`
	Value       string      `json:"value" binding:"required" gorm:"type:varchar;not null"`
}

func (e Contact) IsRDBMSEntity() bool {
	return true
}
