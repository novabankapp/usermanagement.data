package login

import (
	"context"
	"github.com/novabankapp/usermanagement.data/domain/account"
	"github.com/novabankapp/usermanagement.data/domain/login"
)

type AuthRepository interface {
	Create(ctx context.Context, userAccount account.UserAccount, userLogin login.UserLogin) (bool, error)
	Login(ctx context.Context, username string, password string) (bool, error)
	LoginViaOTP(ctx context.Context, otp string) (bool, error)
	ForgotPasswordWithEmail(ctx context.Context, email string) (*string, error)
	ForgotPasswordWithPhone(ctx context.Context, phone string) (*string, error)
}
