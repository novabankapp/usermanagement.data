package auth

import (
	"context"
	"github.com/scylladb/gocqlx/v2"

	"github.com/novabankapp/usermanagement.data/domain/account"
	"github.com/novabankapp/usermanagement.data/domain/login"
)

type AuthRepository interface {
	Create(ctx context.Context, userAccount account.UserAccount, userLogin login.UserLogin) (accountId *string, userId *string, error error)
	VerifyOTP(cxt context.Context, userId string, pin string) (bool, error)
	GetUserById(cxt context.Context, userId string) (*login.UserLogin, error)
	VerifyEmailCode(cxt context.Context, userId string, code string) (bool, error)
	CheckUsername(cxt context.Context, username string) (bool, error)
	CheckEmail(cxt context.Context, email string) (bool, error)
	DeleteUser(cxt context.Context, userId string) (bool, error)
	Login(ctx context.Context, username string, password string) (*[]account.UserAccount, error)
	ChangePassword(ctx context.Context, userId string, oldPassword string, newPassword string) (bool, error)
	IsUserKycCompliant(userId string, ctx context.Context, session *gocqlx.Session) bool
	IsAccountLocked(userId string, ctx context.Context, session *gocqlx.Session) bool
}
