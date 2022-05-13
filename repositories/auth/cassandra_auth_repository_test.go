package auth

import (
	"testing"
	"time"

	"fmt"

	"github.com/fatih/structs"
	"github.com/gocql/gocql"
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
	timeout        time.Duration
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
	s.timeout = 1000
	require.NoError(s.T(), err)
	s.repo = NewCassandraAuthRepository(s.session, s.timeout)
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

	insertAccount := s.userAccountTbl.InsertQuery(*s.session)
	insertAccount.BindStruct(account.UserAccount{})
	batch.Query(insertAccount.String())

	insertUser := s.userLoginTbl.InsertQuery(*s.session)
	insertUser.BindStruct(login.UserLogin{})
	batch.Query(insertUser.String())

	err := s.session.ExecuteBatch(batch)
	require.NoError(s.T(), err)
}
