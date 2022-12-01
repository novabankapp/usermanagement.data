package registration

import "gorm.io/gorm"

type Country struct {
	gorm.Model
	Name string `json:"name" binding:"required" gorm:"type:varchar;not null"`
}

func (e Country) IsRDBMSEntity() bool {
	return true
}
