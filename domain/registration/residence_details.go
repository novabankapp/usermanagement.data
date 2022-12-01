package registration

import "gorm.io/gorm"

type ResidenceDetails struct {
	gorm.Model
	UserID            string  `json:"user_id" binding:"required" gorm:"type:varchar;not null"`
	ResidentialStatus string  `json:"residential_status" binding:"required" gorm:"type:varchar;not null"`
	ProofOfResidency  string  `json:"proof_of_residency" binding:"required" gorm:"type:varchar;not null"`
	NationalityID     int     `gorm:"type:int;not null"`
	Nationality       Country `gorm:"foreignKey:NationalityID" json:"nationality"`
	CountryOfBirthID  int     `gorm:"type:int;not null"`
	CountryOfBirth    Country `gorm:"foreignKey:CountryOfBirthID" json:"country_of_birth"`
}

func (e ResidenceDetails) IsRDBMSEntity() bool {
	return true
}
