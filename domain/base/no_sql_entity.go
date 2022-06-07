package base

import (
	"github.com/novabankapp/usermanagement.data/domain/account"
	"github.com/novabankapp/usermanagement.data/domain/login"
)

type NoSqlEntity interface {
	account.KycCompliant | account.UserAccount | account.UserAccountActivity | login.UserLogin | login.UserLoginAttempts | login.CodeLogin | login.OtpLogin
}
