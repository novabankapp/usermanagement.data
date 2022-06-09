package base

import (
	"github.com/novabankapp/usermanagement.data/domain/account"
	"github.com/novabankapp/usermanagement.data/domain/login"
	"github.com/novabankapp/usermanagement.data/domain/password"
)

type NoSqlEntity interface {
	account.KycCompliant | account.UserAccount | account.UserAccountActivity | login.UserLogin | login.UserLoginAttempts | login.CodeLogin | login.OtpLogin | password.PhonePasswordReset | password.EmailPasswordReset
}
