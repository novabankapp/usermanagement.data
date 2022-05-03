package base

import (
	"github.com/google/uuid"
	"github.com/novabankapp/usermanagement.data/domain"
	"reflect"
)

type Entity interface {
	domain.Contact | domain.ContactType | domain.Country | domain.ResidenceDetails
	domain.User | domain.UserDetails | domain.UserEmployment | domain.UserIdentification | domain.UserIncome
}

func FillDefaults[E Entity](entity E) {
	metaValue := reflect.ValueOf(entity).Elem()
	if metaValue.Type() == reflect.TypeOf("") {
		metaValue.FieldByName("ID").SetString(uuid.New().String())
	}
}
