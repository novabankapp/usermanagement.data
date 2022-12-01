package registration

import "gorm.io/gorm"

type ContactType struct {
	gorm.Model
	Name string `json:"name" binding:"required" gorm:"type:varchar;not null"`
}

func (e ContactType) IsRDBMSEntity() bool {
	return true
}
