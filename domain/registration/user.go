package registration

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
	"time"
)

type UserID struct {
	ID string `uri:"id" binding:"required"`
}

type User struct {
	ID                 string             `gorm:"primaryKey" json:"id" gorm:"type:varchar;not null"`
	FirstName          string             `json:"firstname" binding:"required" gorm:"type:varchar;not null"`
	LastName           string             `json:"lastname" binding:"required" gorm:"type:varchar;not null"`
	UserName           string             `json:"username" binding:"required" gorm:"type:varchar;not null"`
	CreatedAt          time.Time          `json:"created_at" gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt          time.Time          `json:"updated_at"`
	Email              string             `json:"email" gorm:"type:varchar;not null"`
	Phone              string             `json:"phone" gorm:"type:varchar;not null"`
	DeletedAt          gorm.DeletedAt     `json:"deleted_at"`
	UserDetails        UserDetails        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" gorm:"foreignKey:ID"`
	ResidenceDetails   ResidenceDetails   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" gorm:"foreignKey:ID"`
	UserIdentification UserIdentification `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" gorm:"foreignKey:ID"`
	UserIncome         UserIncome         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" gorm:"foreignKey:ID"`
	UserEmployment     UserEmployment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" gorm:"foreignKey:ID"`
	Contacts           []Contact          `json:"contacts" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" gorm:"foreignKey:ID"`
	UserOneTimePin     UserOneTimePin     `json:"one_time_pin" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" gorm:"foreignKey:ID"`
}

func (e User) IsRDBMSEntity() bool {
	return true
}

func (u *User) FillDefaults() {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
}

func (u *User) PrepareCreate() error {
	if u.Email != "" {
		u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	}

	return nil
}
