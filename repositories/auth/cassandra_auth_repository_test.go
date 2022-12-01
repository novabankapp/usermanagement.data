package auth

import (
	"context"
	"testing"
	"time"

	"fmt"

	"github.com/fatih/structs"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/novabankapp/usermanagement.data/constants"
	"github.com/novabankapp/usermanagement.data/domain/account"
	"github.com/novabankapp/usermanagement.data/domain/login"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/gocqlxtest"
	"github.com/scylladb/gocqlx/v2/table"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	session        *gocqlx.Session
	context        context.Context
	repo           AuthRepository
	userLoginTbl   *table.Table
	userAccountTbl *table.Table
	loginOtpTbl    *table.Table
	loginCodeTbl   *table.Table
}

func (s *Suite) SetupSuite() {
	cluster := gocqlxtest.CreateCluster()
	sess, err := gocqlx.WrapSession(cluster.CreateSession())
	s.session = &sess
	require.NoError(s.T(), err)
	error := s.session.ExecStmt(`CREATE KEYSPACE IF NOT EXISTS examples WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}`)
	require.NoError(s.T(), error)
	timeout := time.Duration(10000)
	s.context, _ = context.WithTimeout(context.Background(), timeout)
	require.NoError(s.T(), err)
	s.repo = NewCassandraAuthRepository(s.session, timeout)
	userColumns := structs.Names(&login.UserLogin{})
	userLoginMetadata := table.Metadata{
		Name:    fmt.Sprintf("examples.%s", constants.USERLOGIN),
		Columns: userColumns,
		PartKey: []string{"id"},
	}
	s.userLoginTbl = table.New(userLoginMetadata)

	accountColumns := structs.Names(&account.UserAccount{})
	userAccountMetadata := table.Metadata{
		Name:    fmt.Sprintf("examples.%s", constants.USERACCOUNT),
		Columns: accountColumns,
		PartKey: []string{"id"},
	}
	s.userAccountTbl = table.New(userAccountMetadata)

	loginOtpColumns := structs.Names(&login.OtpLogin{})
	loginOtpMetadata := table.Metadata{
		Name:    fmt.Sprintf("examples.%s", constants.LOGINOTP),
		Columns: loginOtpColumns,
		PartKey: []string{"id"},
	}
	s.loginOtpTbl = table.New(loginOtpMetadata)

	loginCodeColumns := structs.Names(&login.CodeLogin{})
	loginCodeMetadata := table.Metadata{
		Name:    fmt.Sprintf("examples.%s", constants.LOGINCODE),
		Columns: loginCodeColumns,
		PartKey: []string{"id"},
	}
	s.loginCodeTbl = table.New(loginCodeMetadata)

}
func (s *Suite) AfterTest(_, _ string) {
	s.session.Close()
}
func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}
func (s *Suite) Test_repository_Create() {

	ts := time.Now().UnixNano() / 1000
	batch := s.session.NewBatch(gocql.LoggedBatch).WithTimestamp(ts)

	insertAccount := s.userAccountTbl.InsertQuery(*s.session).WithContext(s.context)
	userId := uuid.New().String()
	password, err := login.HashPassword("lmsasajnr@20")
	require.NoError(s.T(), err)
	insertAccount.BindStruct(account.UserAccount{
		ID:        uuid.New().String(),
		UserID:    userId,
		CreatedAt: time.Now(),
		IsLocked:  false,
		IsActive:  true,
	})
	batch.Query(insertAccount.String())

	insertUser := s.userLoginTbl.InsertQuery(*s.session).WithContext(s.context)
	insertUser.BindStruct(login.UserLogin{
		ID:        uuid.New().String(),
		UserID:    userId,
		Email:     "lmsasajnr@gmail.com",
		FirstName: "Lewis",
		LastName:  "Msasa",
		UserName:  "lmsasajnr",
		Password:  *password,
		CreatedAt: time.Now(),
	})
	batch.Query(insertUser.String())

	err = s.session.ExecuteBatch(batch)
	require.NoError(s.T(), err)
}

func (s *Suite) Test_repository_expired_VerifyCode() {
	code := "1234"
	now := time.Now()
	past := now.AddDate(0, 0, -1)
	query := s.session.Query(
		fmt.Sprintf(`INSERT INTO examples.%s (id, user_id,code,expiry_date) VALUES (?, ?, ?, ?)`, constants.LOGINCODE),
		[]string{uuid.New().String(), USERID, code, past.String()},
	)
	err := query.Exec()
	require.NoError(s.T(), err)

	var otp string
	err = s.session.Query(fmt.Sprintf(`SELECT code FROM examples.%s WHERE user_id = ? and code = ? and expiry_date >= ? LIMIT 1`, constants.LOGINCODE),
		[]string{USERID, code, now.String()}).Consistency(gocql.One).WithContext(s.context).Scan(&otp)

	require.NoError(s.T(), err)
}
func (s *Suite) Test_repository_valid_VerifyCode() {
	code := "1234"
	now := time.Now()
	query := s.session.Query(
		fmt.Sprintf(`INSERT INTO examples.%s (id, user_id,code,expiry_date) VALUES (?, ?, ?, ?)`, constants.LOGINCODE),
		[]string{uuid.New().String(), USERID, code, now.Add(time.Duration(10000)).String()},
	)
	err := query.Exec()
	require.NoError(s.T(), err)

	var otp string
	err = s.session.Query(fmt.Sprintf(`SELECT code FROM examples.%s WHERE user_id = ? and code = ? and expiry_date >= ? LIMIT 1`, constants.LOGINCODE),
		[]string{USERID, code, now.String()}).Consistency(gocql.One).WithContext(s.context).Scan(&otp)

	require.NoError(s.T(), err)
}
func (s *Suite) Test_repository_expired_VerifyOTP() {
	pin := "1234"
	now := time.Now()
	past := now.AddDate(0, 0, -1)
	query := s.session.Query(
		fmt.Sprintf(`INSERT INTO examples.%s (id, user_id,pin,expiry_date) VALUES (?, ?, ?, ?)`, constants.LOGINOTP),
		[]string{uuid.New().String(), USERID, pin, past.String()},
	)
	err := query.Exec()
	require.NoError(s.T(), err)

	var otp string
	err = s.session.Query(fmt.Sprintf(`SELECT pin FROM examples.%s WHERE user_id = ? and pin = ? and expiry_date >= ? LIMIT 1`, constants.LOGINOTP),
		[]string{USERID, pin, now.String()}).Consistency(gocql.One).WithContext(s.context).Scan(&otp)

	require.NoError(s.T(), err)
}
func (s *Suite) Test_repository_valid_VerifyOTP() {
	pin := "1234"
	now := time.Now()
	query := s.session.Query(
		fmt.Sprintf(`INSERT INTO examples.%s (id, user_id,pin,expiry_date) VALUES (?, ?, ?, ?)`, constants.LOGINOTP),
		[]string{uuid.New().String(), USERID, pin, now.Add(time.Duration(10000)).String()},
	)
	err := query.Exec()
	require.NoError(s.T(), err)

	var otp string
	err = s.session.Query(fmt.Sprintf(`SELECT pin FROM examples.%s WHERE user_id = ? and pin = ? and expiry_date >= ? LIMIT 1`, constants.LOGINOTP),
		[]string{USERID, pin, now.String()}).Consistency(gocql.One).WithContext(s.context).Scan(&otp)

	require.NoError(s.T(), err)
}

const USERID = "12x1234"

func createUser(session *gocqlx.Session) error {
	query := session.Query(
		fmt.Sprintf(`INSERT INTO examples.%s (id, user_id, email, first_name,last_name,user_name,password,created_at) VALUES (?, ?, ?, ?,?,?,?,?)`, constants.USERLOGIN),
		[]string{uuid.New().String(), USERID, "lmsasajnr@gmail.com", "Lewis", "Msasa", "lmsasajnr", "123432", time.Now().String()},
	)

	if err := query.Exec(); err != nil {
		return err
	}
	return nil
}
