package login

import (
	"github.com/gocql/gocql"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserLogin struct {
	ID        gocql.UUID `json:"id"`
	UserID    string     `json:"user_id" binding:"required"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	UserName  string     `json:"user_name" binding:"required"`
	Password  string     `json:"password"`
	Pin       string     `json:"pin"`
	IsActive  bool       `json:"is_active"`
	IsLocked  bool       `json:"is_locked"`
	CreatedAt time.Time  `json:"created_at"`
}

func (k UserLogin) IsNoSQLEntity() bool {
	//gocql.RandomUUID()
	return true
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
