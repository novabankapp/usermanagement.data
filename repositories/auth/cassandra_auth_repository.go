package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/fatih/structs"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/novabankapp/usermanagement.data/constants"
	"github.com/novabankapp/usermanagement.data/domain/account"
	"github.com/novabankapp/usermanagement.data/domain/login"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
)

type CassandraAuthRepository struct {
	session *gocqlx.Session
	timeout time.Duration
}

func NewCassandraAuthRepository(session *gocqlx.Session, timeout time.Duration) AuthRepository {
	return &CassandraAuthRepository{
		session: session,
		timeout: timeout,
	}
}
func (repo CassandraAuthRepository) VerifyOTP(cxt context.Context, userId string, pin string) (bool, error) {
	ctx, cancel := context.WithTimeout(cxt, repo.timeout)
	defer cancel()
	now := time.Now()
	var otp string
	if err := repo.session.Query(fmt.Sprintf(`SELECT pin FROM %s WHERE user_id = ? and pin = ? and expiry_date >= ? LIMIT 1`, constants.LOGINOTP),
		[]string{userId, pin, now.String()}).Consistency(gocql.One).WithContext(ctx).Scan(&otp); err != nil {
		return false, err
	}
	return true, nil
}
func (repo CassandraAuthRepository) VerifyEmailCode(cxt context.Context, userId string, code string) (bool, error) {
	ctx, cancel := context.WithTimeout(cxt, repo.timeout)
	defer cancel()
	now := time.Now()
	if err := repo.session.Query(fmt.Sprintf(`SELECT code FROM %s WHERE user_id = ? and code = ? and expiry_date >= ? LIMIT 1`, constants.LOGINCODE),
		[]string{userId, code, now.String()}).Consistency(gocql.One).WithContext(ctx).Scan(&code); err != nil {
		return false, err
	}
	return true, nil

}
func (repo CassandraAuthRepository) ForgotPasswordWithEmail(ctx context.Context, email string) (*string, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.timeout)
	defer cancel()
	var userId string
	if err := repo.session.Query(fmt.Sprintf(`SELECT user_id FROM %s WHERE email = ? LIMIT 1`, constants.USERLOGIN),
		[]string{email}).Consistency(gocql.One).WithContext(ctx).Scan(&userId); err != nil {
		return nil, err
	}
	return &userId, nil
}
func (repo CassandraAuthRepository) ForgotPasswordWithPhone(ctx context.Context, phone string) (*string, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.timeout)
	defer cancel()
	var userId string
	if err := repo.session.Query(fmt.Sprintf(`SELECT user_id FROM %s WHERE phone = ? LIMIT 1`, constants.USERLOGIN),
		[]string{phone}).Consistency(gocql.One).WithContext(ctx).Scan(&userId); err != nil {
		return nil, err
	}
	return &userId, nil
}
func (repo CassandraAuthRepository) ChangePassword(ctx context.Context, userId string, oldPassword string, newPassword string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.timeout)
	defer cancel()
	passwordHashed, err := login.HashPassword(newPassword)
	if err != nil {
		return false, err
	}
	oldPasswordHashed, err := login.HashPassword(oldPassword)
	if err != nil {
		return false, err
	}
	q := qb.Update(constants.USERLOGIN).
		SetLit("password", *passwordHashed).
		Where(qb.EqLit("user_id", userId)).
		Where(qb.EqLit("password", *oldPasswordHashed)).
		Query(*repo.session).
		SerialConsistency(gocql.Serial).
		WithContext(ctx)

	applied, err := q.ExecCASRelease()
	if err != nil {
		return false, err
	}
	return applied, nil
}
func (repo CassandraAuthRepository) Login(ctx context.Context, username string, password string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.timeout)
	defer cancel()

	hashed, err := login.HashPassword(password)
	if err != nil {
		return false, err
	}
	var userId string
	if err := repo.session.Query(fmt.Sprintf(`SELECT user_id FROM %s WHERE username = ? and password = ? LIMIT 1`, constants.USERLOGIN),
		[]string{username, *hashed}).Consistency(gocql.One).WithContext(ctx).Scan(&userId); err != nil {
		return false, err
	}

	getUserAccount := qb.Select(constants.USERACCOUNT).
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
func IsAccountLocked(userId string, ctx context.Context, session *gocqlx.Session) bool {
	getUserAccount := qb.Select(constants.USERACCOUNT).
		Where(qb.EqLit("user_id", userId)).
		Query(*session).
		WithContext(ctx)
	var results []*account.UserAccount
	errr := getUserAccount.Select(&results)
	if errr != nil || len(results) < 1 {
		return false
	}
	acc := results[0]
	if acc.IsLocked {
		return false
	}
	if !acc.IsActive {
		return false
	}
	return true
}
func (repo CassandraAuthRepository) Create(ctx context.Context, userAccount account.UserAccount,
	userLogin login.UserLogin) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.timeout)
	defer cancel()

	accountColumns := structs.Names(&account.UserAccount{})
	ts := time.Now().UnixNano() / 1000
	batch := repo.session.NewBatch(gocql.LoggedBatch).WithTimestamp(ts)
	userAccount.ID = uuid.New().String()
	userAccount.CreatedAt = time.Now()
	userAccount.IsActive = true
	userAccount.IsLocked = false
	userAccount.IsKyc = false
	insertAccount := qb.Insert(constants.USERACCOUNT).
		Columns(accountColumns...).
		Query(*repo.session).
		WithContext(ctx)
	insertAccount.BindStruct(userAccount)
	batch.Query(insertAccount.String())

	userLogin.ID = uuid.New().String()
	userLogin.CreatedAt = time.Now()
	userLogin.HashPassword()
	userLoginColumns := structs.Names(&login.UserLogin{})
	insertUserLogin := qb.Insert(constants.USERLOGIN).
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
