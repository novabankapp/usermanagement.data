package cassandra

import (
	"context"
	"errors"
	"github.com/fatih/structs"
	"github.com/gocql/gocql"
	"github.com/novabankapp/usermanagement.data/repositories/base"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"log"
	"time"
)

type CassandraRepository struct {
	session   *gocqlx.Session
	tableName string
	timeout   time.Duration
}

const timeout time.Duration = 10000

func NewCassandraRepository(session *gocqlx.Session, tableName string, timeout time.Duration) base.Repository {
	return &CassandraRepository{
		session:   session,
		tableName: tableName,
		timeout:   timeout,
	}
}
func GetById[E base.Entity](rep *CassandraRepository, ctx context.Context, id string) (*E, error) {

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var result []E
	getUser := qb.Select(rep.tableName).
		Where(qb.EqLit("id", id)).
		Query(*rep.session).
		WithContext(ctx)

	err := getUser.Select(&result)
	if err != nil {
		return nil, err
	}
	if len(result) < 1 {
		return nil, errors.New("record not found")
	}
	return &result[0], nil
}
func Create[E base.Entity](rep *CassandraRepository, ctx context.Context, entity E) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	columns := structs.Names(&E{})
	insert := qb.Insert(rep.tableName).
		Columns(columns...).
		Query(*rep.session).
		WithContext(ctx)
	insert.BindStruct(entity)
	if err := insert.ExecRelease(); err != nil {
		log.Println(err)
		return false, err
	}
	return true, nil
}
func Update[E base.Entity](rep *CassandraRepository, ctx context.Context, entity E, id string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	columns := structs.Names(&E{})
	updateUser := qb.Update(rep.tableName).
		Set(columns...).
		Where(qb.EqLit("id", id)).
		Query(*rep.session).
		SerialConsistency(gocql.Serial).WithContext(ctx)

	updateUser.BindStruct(entity)

	applied, err := updateUser.ExecCASRelease()
	if err != nil {
		return false, err
	}

	return applied, nil
}
func Delete[E base.Entity](rep *CassandraRepository, ctx context.Context, id string) (bool, error) {

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	ent, error := GetById(rep, ctx, id)

	if error != nil {
		return false, error
	}
	delete := qb.Delete(rep.tableName).
		Where(qb.EqLit("id", id)).
		Query(*rep.session).
		SerialConsistency(gocql.Serial).WithContext(ctx)

	delete.BindStruct(ent)
	applied, err := delete.ExecCASRelease()
	if err != nil {
		return false, err
	}

	return applied, nil
}

func Get[E base.Entity](rep *CassandraRepository, ctx context.Context,
	page []byte, pageSize int, query string, orderBy string) (*[]E, error) {

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var results []E
	get := qb.Select(rep.tableName).
		Query(*rep.session).
		PageSize(pageSize).
		PageState(page).
		WithContext(ctx)

	err := get.Select(&results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}
