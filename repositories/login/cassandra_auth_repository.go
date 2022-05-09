package login

import (
	"context"
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
}

func (repo CassandraAuthRepository) Login(ctx context.Context, username string, password string) (bool, error) {
	var results []login.UserLogin
	hashed, err := login.HashPassword(password)
	if err != nil {
		return false, nil
	}
	getUser := qb.Select("user_login").
		Where(qb.EqLit("username", username), qb.EqLit("password", *hashed)).
		Query(*repo.session).
		WithContext(ctx)

	errr := getUser.Select(&results)
	if errr != nil {
		return false, errr
	}
	return true, nil
}
func (repo CassandraAuthRepository) Create(ctx context.Context, userAccount account.UserAccount, userLogin login.UserLogin) (bool, error) {
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
	insertUserLogin.BindStruct(insertUserLogin)
	batch.Query(insertUserLogin.String())
	if err := repo.session.ExecuteBatch(batch); err != nil {
		return false, err
	}

	return true, nil
}
