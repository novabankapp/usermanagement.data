package domain

type ResidenceDetails struct {
	UserID            string `json:"user_id" binding:"required"`
	ResidentialStatus string `json:"residential_status" binding:"required"`
	ProofOfResidency  string `json:"proof_of_residency" binding:"required"`
	NationalityID     int
	Nationality       Country `gorm:"foreignKey:NationalityID" json:"nationality"`
	CountryOfBirthID  int
	CountryOfBirth    Country `gorm:"foreignKey:CountryOfBirthID" json:"country_of_birth"`
}
