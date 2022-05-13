package registration

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/novabankapp/usermanagement.data/domain/registration"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	createUserQuery = `INSERT INTO users (id,user_name,first_name, last_name, email, password) 
		VALUES ($1, $2, $3, $4, $5,$6)) 
		RETURNING id`

	findByEmailQuery = `SELECT user_id, email, first_name, last_name, role, avatar, password, created_at, updated_at FROM users WHERE email = $1`

	findByIDQuery = `SELECT user_id, email, first_name, last_name, role, avatar, created_at, updated_at FROM users WHERE user_id = $1`
)

type Suite struct {
	suite.Suite
	DB         *gorm.DB
	mock       sqlmock.Sqlmock
	repository RegisterRepository
	user       *registration.User
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	require.NoError(s.T(), err)

	s.repository = NewPostgresRegisterRepository(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}
func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_repository_Create() {
	id := uuid.New().String()
	user := registration.User{
		ID:        id,
		UserName:  "lmsasajnr",
		Email:     "email@gmail.com",
		FirstName: "FirstName",
		LastName:  "LastName",
		Password:  "123456",
	}
	s.mock.ExpectQuery(regexp.QuoteMeta(createUserQuery)).
		WithArgs(id, user.UserName, user.FirstName, user.LastName, user.Email, user.Password).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(id))

	_, err := s.repository.Create(context.Background(), user)

	require.NoError(s.T(), err)
}

/*func TestRegisterRepository_Create(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()
	conn, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	registerPGRepository := NewPostgresRegisterRepository(conn)

	columns := []string{"user_id", "first_name", "last_name", "email", "password", "avatar", "role", "created_at", "updated_at"}
	userUUID := uuid.New()
	mockUser := &registration.User{
		ID:        userUUID.String(),
		Email:     "email@gmail.com",
		FirstName: "FirstName",
		LastName:  "LastName",
		Password:  "123456",
	}

	rows := sqlmock.NewRows(columns).AddRow(
		userUUID,
		mockUser.FirstName,
		mockUser.LastName,
		mockUser.Email,
		mockUser.Password,
		time.Now(),
		time.Now(),
	)

	mock.ExpectQuery(createUserQuery).WithArgs(
		mockUser.FirstName,
		mockUser.LastName,
		mockUser.Email,
		mockUser.Password,
	).WillReturnRows(rows)

	createdUser, err := registerPGRepository.Create(context.Background(), *mockUser)
	require.NoError(t, err)
	require.NotNil(t, createdUser)
}*/
