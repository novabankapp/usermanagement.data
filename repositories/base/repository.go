package base

import (
	"context"
	domainbase "github.com/novabankapp/usermanagement.data/domain/base"
)

type Repository[E domainbase.Entity] interface {
	Create(ctx context.Context, entity E) (*E, error)
	Update(ctx context.Context, entity E) (bool, error)
	Delete(ctx context.Context, ID string) (bool, error)
	GetById(ctx context.Context, ID string) (*E, error)
	Get(ctx context.Context, page int, pageSize int, query string, orderBy string) (*[]E, error)
}
