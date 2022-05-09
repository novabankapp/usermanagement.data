package auth

import (
	"context"
	"errors"
	"github.com/fatih/structs"
	"github.com/gocql/gocql"
	"github.com/novabankapp/usermanagement.data/domain/account"
	"github.com/novabankapp/usermanagement.data/domain/login"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"time"
)

type CassandraAuthRepository struct {
	session *gocqlx.Session
	timeout time.Duration
}

func (repo CassandraAuthRepository) Login(ctx context.Context, username string, password string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.timeout)
	defer cancel()

	hashed, err := login.HashPassword(password)
	if err != nil {
		return false, err
	}
	var userId string
	if err := repo.session.Query(`SELECT user_id FROM user_login WHERE username = ? and password = ? LIMIT 1`,
		[]string{username, *hashed}).Consistency(gocql.One).Scan(&userId); err != nil {
		return false, err
	}

	getUserAccount := qb.Select("user_account").
		Where(qb.EqLit("user_id", userId)).
		Query(*repo.session).
		WithContext(ctx)
	var results []*account.UserAccount
	errr := getUserAccount.Select(&results)
	if errr != nil || len(results) < 1 {
		return false, errr
	}
	acc := results[0]
	if acc.IsLocked {
		return false, errors.New("account locked")
	}
	if !acc.IsActive {
		return false, errors.New("account not active")
	}
	return true, nil
}
func (repo CassandraAuthRepository) Create(ctx context.Context, userAccount account.UserAccount, userLogin login.UserLogin) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.timeout)
	defer cancel()

	accountColumns := structs.Names(&account.UserAccount{})
	ts := time.Now().UnixNano() / 1000
	batch := repo.session.NewBatch(gocql.LoggedBatch).WithTimestamp(ts)
	insertAccount := qb.Insert("user_account").
		Columns(accountColumns...).
		Query(*repo.session).
		WithContext(ctx)
	insertAccount.BindStruct(userAccount)
	batch.Query(insertAccount.String())

	userLoginColumns := structs.Names(&login.UserLogin{})
	insertUserLogin := qb.Insert("user_login").
		Columns(userLoginColumns...).
		Query(*repo.session).
		WithContext(ctx)
	insertUserLogin.BindStruct(userLogin)
	batch.Query(insertUserLogin.String())
	if err := repo.session.ExecuteBatch(batch); err != nil {
		return false, err
	}

	return true, nil
}
