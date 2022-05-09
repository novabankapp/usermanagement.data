package account

import "context"

type AccountRepository interface {
	LockAccount(ctx context.Context, userId string) (bool, error)
	UnlockAccount(ctx context.Context, userId string) (bool, error)
	DeactivateAccount(ctx context.Context, userId string) (bool, error)
	ActivateAccount(ctx context.Context, userId string) (bool, error)
	LogAccountActivity(ctx context.Context, activity string, userId string) (bool, error)
}
