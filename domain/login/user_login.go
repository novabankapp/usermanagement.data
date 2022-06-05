package login

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserLogin struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id" binding:"required"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	UserName  string    `json:"username" binding:"required"`
	Password  string    `json:"password"`
	Pin       string    `json:"pin"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *UserLogin) HashPassword() error {
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	if u.Pin != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Pin), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Pin = string(hashedPassword)
	}

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
func (u *UserLogin) ComparePasswords(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
