package init

import (
	"github.com/novabankapp/common.infrastructure/postgres"
	"github.com/novabankapp/usermanagement.data/domain/registration"
)

func Init(config postgres.Config) {
	postgres.AutoMigrateDB(config, registration.User{}, registration.Contact{},
		registration.ContactType{}, registration.UserEmployment{},
		registration.UserDetails{}, registration.UserIncome{},
		registration.UserIdentification{}, registration.Country{},
		registration.ResidenceDetails{}, registration.UserOneTimePin{})
}
