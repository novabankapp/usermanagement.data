package init

import (
	"context"

	"github.com/novabankapp/common.infrastructure/postgres"
	"github.com/novabankapp/usermanagement.data/domain/registration"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/migrate"
)

func InitPostgres(config postgres.Config) {
	postgres.AutoMigrateDB(config, registration.User{}, registration.Contact{},
		registration.ContactType{}, registration.UserEmployment{},
		registration.UserDetails{}, registration.UserIncome{},
		registration.UserIdentification{}, registration.Country{},
		registration.ResidenceDetails{}, registration.UserOneTimePin{})
}
func InitCassandra(session *gocqlx.Session) {
	log := func(ctx context.Context, session gocqlx.Session, ev migrate.CallbackEvent, name string) error {
		return nil
	}
	reg := migrate.CallbackRegister{}
	reg.Add(migrate.BeforeMigration, "cassandra_migrate.cql", log)
	reg.Add(migrate.AfterMigration, "cassandra_migrate.cql", log)
	reg.Add(migrate.CallComment, "1", log)
	reg.Add(migrate.CallComment, "2", log)
	reg.Add(migrate.CallComment, "3", log)
	migrate.Callback = reg.Callback

	// First run prints data
	if err := migrate.FromFS(context.Background(), *session, Files); err != nil {

	}

	// Second run skips the processed files
	if err := migrate.FromFS(context.Background(), *session, Files); err != nil {

	}
}
