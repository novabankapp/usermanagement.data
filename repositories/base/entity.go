package base

import (
	"github.com/google/uuid"
	"github.com/novabankapp/usermanagement.data/domain/registration"
	"reflect"
)

type Entity interface {
	registration.Contact | registration.ContactType | registration.Country | registration.ResidenceDetails
	registration.User | registration.UserDetails | registration.UserEmployment | registration.UserIdentification | registration.UserIncome
}

func FillDefaults[E Entity](entity E) {
	metaValue := reflect.ValueOf(entity).Elem()
	if metaValue.Type() == reflect.TypeOf("") {
		metaValue.FieldByName("ID").SetString(uuid.New().String())
	}
}
