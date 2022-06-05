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
	ID                 string             `gorm:"primaryKey" json:"id"`
	FirstName          string             `json:"firstname" binding:"required"`
	LastName           string             `json:"lastname" binding:"required"`
	UserName           string             `json:"username" binding:"required"`
	CreatedAt          time.Time          `json:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at"`
	Email              string             `json:"email"`
	Phone              string             `json:"phone"`
	DeletedAt          gorm.DeletedAt     `json:"deleted_at"`
	UserDetails        UserDetails        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ResidenceDetails   ResidenceDetails   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserIdentification UserIdentification `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserIncome         UserIncome         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserEmployment     UserEmployment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Contacts           []Contact          `json:"contacts" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserOneTimePin     UserOneTimePin     `json:"one_time_pin" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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
