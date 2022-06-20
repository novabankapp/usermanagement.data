package registration

import "gorm.io/gorm"

type Country struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
}

func (e Country) IsRDBMSEntity() bool {
	return true
}
