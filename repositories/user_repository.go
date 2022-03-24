package repositories

import (
	"context"
	"github.com/novabankapp/usermanagement/usermanagement.data/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user domain.User) (*string, error)
	Update(ctx context.Context, user domain.User) (bool, error)
	Delete(ctx context.Context, user domain.User) (bool, error)
	GetUser(ctx context.Context, ID string) (*domain.User, error)
	GetUsers(ctx context.Context, page int, pageSize int, query string, orderBy string) (*[]domain.User, error)
}
