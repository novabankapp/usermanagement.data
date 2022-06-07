package account

import (
	"github.com/fatih/structs"
	"reflect"
)

type KycCompliant struct {
	UserId                string `json:"user_id"`
	HasUserDetails        bool   `json:"has_user_details"`
	HasResidenceDetails   bool   `json:"has_residence_details"`
	HasUserIdentification bool   `json:"has_user_identification"`
	HasUserIncome         bool   `json:"has_user_income"`
	HasUserEmployment     bool   `json:"has_user_employment"`
}

func (k *KycCompliant) IsKycCompliant() bool {
	fields := structs.Fields(k)
	compliant := false
	for _, field := range fields {
		if field.Name() == "UserId" {
			continue
		}

		val := field.Value()
		if field.Kind() == reflect.Bool {
			elemValue := reflect.ValueOf(val)
			if elemValue.Bool() == false {
				return false
			} else {
				compliant = true
			}
		}

	}
	return compliant

}
