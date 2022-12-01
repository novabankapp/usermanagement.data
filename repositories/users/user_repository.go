package users

import (
	"context"

	"github.com/novabankapp/usermanagement.data/domain/registration"
)

type UserRepository interface {
	Create(ctx context.Context, user registration.User) (*string, error)
	Update(ctx context.Context, user registration.User) (bool, error)
	Delete(ctx context.Context, user registration.User) (bool, error)
	GetUser(ctx context.Context, ID string) (*registration.User, error)
	GetUsers(ctx context.Context, page int, pageSize int, query *string, orderBy *string) (*[]registration.User, error)
}
