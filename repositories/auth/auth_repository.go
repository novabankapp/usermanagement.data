package auth

import (
	"context"

	"github.com/novabankapp/usermanagement.data/domain/account"
	"github.com/novabankapp/usermanagement.data/domain/login"
)

type AuthRepository interface {
	Create(ctx context.Context, userAccount account.UserAccount, userLogin login.UserLogin) (bool, error)
	VerifyOTP(cxt context.Context, userId string, pin string) (bool, error)
	VerifyEmailCode(cxt context.Context, userId string, code string) (bool, error)
	Login(ctx context.Context, username string, password string) (bool, error)
	ForgotPasswordWithEmail(ctx context.Context, email string) (*string, error)
	ForgotPasswordWithPhone(ctx context.Context, phone string) (*string, error)
	ChangePassword(ctx context.Context, userId string, oldPassword string, newPassword string) (bool, error)
}
