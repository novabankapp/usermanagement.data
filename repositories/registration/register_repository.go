package registration

import (
	"context"
	"github.com/novabankapp/usermanagement.data/domain/registration"
)

type RegisterRepository interface {
	Create(ctx context.Context, user registration.User) (*string, error)
	VerifyEmail(ctx context.Context, email string, code string) (bool, error)
	VerifyPhone(ctx context.Context, email string, code string) (bool, error)
	VerifyOTP(cxt context.Context, userId string, pin string) (bool, error)
	SaveDetails(cxt context.Context, userId string, details registration.UserDetails) (bool, error)
	SaveResidenceDetails(cxt context.Context, userId string, details registration.ResidenceDetails) (bool, error)
	SaveUserIdentification(cxt context.Context, userId string, identification registration.UserIdentification) (bool, error)
	SaveUserIncome(cxt context.Context, userId string, income registration.UserIncome) (bool, error)
	SaveEmployment(cxt context.Context, userId string, employment registration.UserEmployment) (bool, error)
	SaveContact(cxt context.Context, userId string, contact registration.Contact) (bool, error)
}

