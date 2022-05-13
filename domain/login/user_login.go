package login

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserLogin struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Phone     string    `json:"phone" binding:"required"`
	FirstName string    `json:"firstname" binding:"required"`
	LastName  string    `json:"lastname" binding:"required"`
	UserName  string    `json:"username" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *UserLogin) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
func HashPassword(password string) (*string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	pass := string(hashedPassword)
	return &pass, nil
}
