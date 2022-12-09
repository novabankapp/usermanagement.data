package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/fatih/structs"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/novabankapp/common.data/utils"
	"github.com/novabankapp/usermanagement.data/constants"
	"github.com/novabankapp/usermanagement.data/domain/account"
	"github.com/novabankapp/usermanagement.data/domain/login"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
)

type cassandraAuthRepository struct {
	session *gocqlx.Session
	timeout time.Duration
}

func NewCassandraAuthRepository(session *gocqlx.Session, timeout time.Duration) AuthRepository {
	return &cassandraAuthRepository{
		session: session,
		timeout: timeout,
	}
}

func (repo cassandraAuthRepository) CheckUsername(cxt context.Context, username string) (bool, error) {
	ctx, cancel := context.WithTimeout(cxt, repo.timeout)
	defer cancel()
	var userId string
	if err := repo.session.Query(fmt.Sprintf(`SELECT user_id FROM %s WHERE username = ? LIMIT 1`, constants.USERLOGIN),
		[]string{username}).Consistency(gocql.One).WithContext(ctx).Scan(&userId); err != nil {
		return false, err
	}
	return &userId != nil || userId == "", nil
}

func (repo cassandraAuthRepository) CheckEmail(cxt context.Context, email string) (bool, error) {
	ctx, cancel := context.WithTimeout(cxt, repo.timeout)
	defer cancel()
	var userId string
	if err := repo.session.Query(fmt.Sprintf(`SELECT user_id FROM %s WHERE email = ? LIMIT 1`, constants.USERLOGIN),
		[]string{email}).Consistency(gocql.One).WithContext(ctx).Scan(&userId); err != nil {
		return false, err
	}
	return &userId != nil || userId == "", nil
}

func (repo cassandraAuthRepository) VerifyOTP(cxt context.Context, userId string, pin string) (bool, error) {
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
func (repo cassandraAuthRepository) VerifyEmailCode(cxt context.Context, userId string, code string) (bool, error) {
	ctx, cancel := context.WithTimeout(cxt, repo.timeout)
	defer cancel()
	now := time.Now()
	if err := repo.session.Query(fmt.Sprintf(`SELECT code FROM %s WHERE user_id = ? and code = ? and expiry_date >= ? LIMIT 1`, constants.LOGINCODE),
		[]string{userId, code, now.String()}).Consistency(gocql.One).WithContext(ctx).Scan(&code); err != nil {
		return false, err
	}
	return true, nil

}

func (repo cassandraAuthRepository) ChangePassword(ctx context.Context, userId string, oldPassword string, newPassword string) (bool, error) {
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
func (repo cassandraAuthRepository) DeleteUser(cxt context.Context, userId string) (bool, error) {
	var tables = []string{constants.USERLOGIN, constants.USERACCOUNT, constants.KYC_COMPLIANT, constants.USERACTIVITIES}

	ts := time.Now().UnixNano() / 1000
	batch := repo.session.NewBatch(gocql.LoggedBatch).WithTimestamp(ts)
	for i := range tables {
		delete := qb.Delete(tables[i]).
			Where(qb.EqLit("user_id", userId)).
			Query(*repo.session).
			SerialConsistency(gocql.Serial).WithContext(cxt)

		batch.Query(delete.String())

	}
	if err := repo.session.ExecuteBatch(batch); err != nil {
		return false, err
	}
	return true, nil

}
func (repo cassandraAuthRepository) GetUserById(cxt context.Context, userId string) (*login.UserLogin, error) {
	cxt, cancel := context.WithTimeout(cxt, repo.timeout)
	defer cancel()
	var user login.UserLogin
	if err := repo.session.Query(fmt.Sprintf(`SELECT * FROM %s WHERE user_id = ? LIMIT 1`, constants.USERLOGIN),
		[]string{userId}).Consistency(gocql.One).WithContext(cxt).Scan(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
func (repo cassandraAuthRepository) Login(ctx context.Context, username string, password string) (*[]account.UserAccount, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.timeout)
	defer cancel()

	hashed, err := login.HashPassword(password)
	if err != nil {
		return nil, err
	}
	var userId string
	var realPassword string
	var active bool
	var locked bool
	if err := repo.session.Query(fmt.Sprintf(`SELECT user_id, password, is_active, is_locked FROM %s WHERE username = ? LIMIT 1`, constants.USERLOGIN),
		[]string{username}).Consistency(gocql.One).WithContext(ctx).Scan(&userId, realPassword, active, locked); err != nil {
		return nil, err
	}
	if !active {
		return nil, errors.New("user is not active")
	}
	if locked {
		return nil, errors.New("account is locked")
	}
	user := login.UserLogin{UserID: userId, Password: realPassword}
	e := user.ComparePasswords(*hashed)
	if e != nil {
		return nil, errors.New("password doesn't match")
	}
	getUserAccounts := qb.Select(constants.USERACCOUNT).
		Where(qb.EqLit("user_id", userId)).
		Query(*repo.session).
		WithContext(ctx)
	var results *[]account.UserAccount
	errr := getUserAccounts.Select(&results)
	if errr != nil {
		return nil, errr
	}
	return results, nil
}
func (repo cassandraAuthRepository) IsAccountLocked(userId string, ctx context.Context, session *gocqlx.Session) bool {
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
func (repo cassandraAuthRepository) IsUserKycCompliant(userId string, ctx context.Context, session *gocqlx.Session) bool {
	getUserAccount := qb.Select(constants.KYC_COMPLIANT).
		Where(qb.EqLit("user_id", userId)).
		Query(*session).
		WithContext(ctx)
	var results []*account.KycCompliant
	error := getUserAccount.Select(&results)
	if error != nil || len(results) < 1 {
		return false
	}
	acc := results[0]

	return acc.IsKycCompliant()
}
func (repo cassandraAuthRepository) Create(ctx context.Context, userAccount account.UserAccount,
	userLogin login.UserLogin) (accountId *string, userId *string, error error) {
	//ctx, cancel := context.WithTimeout(ctx, repo.timeout)
	//defer cancel()

	accountColumns := utils.Map(structs.Names(userAccount), utils.ToSnakeCase)

	//ts := time.Now().UnixNano() / 1000
	//batch := repo.session.NewBatch(gocql.LoggedBatch).WithTimestamp(ts)
	accountGuid, e := gocql.ParseUUID(uuid.New().String())
	if e != nil {
		return nil, nil, e
	}
	userAccount.ID = accountGuid
	userAccount.CreatedAt = time.Now()
	userAccount.IsActive = true
	userAccount.IsLocked = false
	insertAccount := qb.Insert(constants.USERACCOUNT).
		Columns(accountColumns...).
		Query(*repo.session)
	insertAccount.BindStruct(userAccount)
	//batch.Query(insertAccount.Query.String())
	if err := insertAccount.WithContext(ctx).ExecRelease(); err != nil {
		return nil, nil, err
	}
	id, err := gocql.ParseUUID(uuid.New().String())
	if err != nil {

		return nil, nil, err
	}
	userLogin.ID = id
	userLogin.CreatedAt = time.Now()
	userLogin.HashPassword()
	userLoginColumns := utils.Map(structs.Names(userLogin), utils.ToSnakeCase)
	insertUserLogin := qb.Insert(constants.USERLOGIN).
		Columns(userLoginColumns...).
		Query(*repo.session)
	insertUserLogin.BindStruct(userLogin)
	//batch.Query(insertUserLogin.Statement())
	if err := insertUserLogin.WithContext(ctx).ExecRelease(); err != nil {
		fmt.Println(fmt.Sprintf("error creating user login:%s", err.Error()))
		return nil, nil, err
	}
	i := userAccount.ID.String()
	u := userLogin.ID.String()
	return &i, &u, nil
}
