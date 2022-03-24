package repositories

import (
	"context"
	"errors"
	"github.com/gocql/gocql"
	"github.com/novabankapp/usermanagement/usermanagement.data/domain"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"log"
)

type cassandraUserRepository struct {
	session   *gocqlx.Session
	tableName string
}

var columns = make([]string, 5)

func (repo *cassandraUserRepository) GetUsers(ctx context.Context, page int, pageSize int, query string, orderBy string) (*[]domain.User, error) {
	var results []domain.User
	getUser := qb.Select(repo.tableName).
		Where(qb.Eq("id")).
		Query(*repo.session).
		PageSize(pageSize).
		WithContext(ctx)

	err := getUser.Select(&results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}
func (repo *cassandraUserRepository) Create(ctx context.Context, user domain.User) (*string, error) {

	insertUser := qb.Insert(repo.tableName).
		Columns(columns...).
		Query(*repo.session).
		WithContext(ctx)
	insertUser.BindStruct(user)
	if err := insertUser.ExecRelease(); err != nil {
		log.Println(err)
		return nil, err
	}
	result := ""
	return &result, nil
}
func (repo *cassandraUserRepository) Delete(ctx context.Context, user domain.User) (bool, error) {
	deleteUser := qb.Delete(repo.tableName).
		Where(qb.Eq("id")).
		Query(*repo.session).
		SerialConsistency(gocql.Serial).WithContext(ctx)

	deleteUser.BindStruct(user)
	applied, err := deleteUser.ExecCASRelease()
	if err != nil {
		return false, err
	}

	return applied, nil
}
func (repo *cassandraUserRepository) GetUser(ctx context.Context, ID string) (*domain.User, error) {
	var userResult []domain.User
	getUser := qb.Select(repo.tableName).
		Where(qb.Eq("id")).
		Query(*repo.session).
		WithContext(ctx)

	err := getUser.Select(&userResult)
	if err != nil {
		return nil, err
	}
	if len(userResult) < 1 {
		return nil, errors.New("Record not found")
	}
	return &userResult[0], nil
}
func (repo *cassandraUserRepository) Update(ctx context.Context, user domain.User) (bool, error) {
	updateUser := qb.Update(repo.tableName).
		Set(columns...).
		Where(qb.Eq("id")).
		Query(*repo.session).
		SerialConsistency(gocql.Serial).WithContext(ctx)

	updateUser.BindStruct(user)

	applied, err := updateUser.ExecCASRelease()
	if err != nil {
		return false, err
	}

	return applied, nil
}
func NewUserRepositoryCassandra(sess *gocqlx.Session, tableName string) UserRepository {

	return &cassandraUserRepository{
		sess,
		tableName,
	}
}
