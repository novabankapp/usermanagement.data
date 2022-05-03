package init

import (
	"github.com/novabankapp/common.infrastructure/postgres"
	"github.com/novabankapp/usermanagement.data/domain"
)

func Init(config postgres.Config) {
	postgres.AutoMigrateDB(config, domain.User{}, domain.Contact{},
		domain.ContactType{}, domain.UserEmployment{},
		domain.UserDetails{}, domain.UserIncome{},
		domain.UserIdentification{}, domain.Country{},
		domain.ResidenceDetails{})
}
