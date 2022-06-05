package registration

import (
	"context"
	"errors"
	"github.com/novabankapp/usermanagement.data/constants"
	"github.com/novabankapp/usermanagement.data/domain/registration"
	"gorm.io/gorm"
	"time"
)

type postgresRegisterRepository struct {
	conn *gorm.DB
}

func (p postgresRegisterRepository) Create(ctx context.Context, user registration.User) (*string, error) {
	user.FillDefaults()
	err := user.PrepareCreate()
	if err != nil {
		return nil, err
	}
	result := p.conn.Create(&user).WithContext(ctx)
	if result.Error != nil && result.RowsAffected != 1 {
		return nil, errors.New("error occurred while creating a new user")
	}

	return &user.ID, nil
}

func (p postgresRegisterRepository) VerifyEmail(ctx context.Context, email string, code string) (bool, error) {
	var emailActivationCode registration.EmailVerificationCode
	result := p.conn.Model(registration.EmailVerificationCode{
		Email: email,
		Code:  code,
	}).First(&emailActivationCode)
	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, errors.New("code not found")

		} else {
			return false, errors.New("error occurred while fetching code")

		}
		return false, result.Error
	}
	today := time.Now()
	if emailActivationCode.ExpiryDate.Before(today) {
		return false, errors.New("code expired")
	}
	return true, nil
}

func (p postgresRegisterRepository) VerifyPhone(ctx context.Context, phone string, code string) (bool, error) {
	var phoneActivationCode registration.PhoneVerificationCode
	result := p.conn.Model(registration.PhoneVerificationCode{
		Phone: phone,
		Code:  code,
	}).First(&phoneActivationCode)
	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, errors.New("code not found")

		} else {
			return false, errors.New("error occurred while fetching code")

		}
		return false, result.Error
	}
	today := time.Now()
	if phoneActivationCode.ExpiryDate.Before(today) {
		return false, errors.New("code expired")
	}
	return true, nil
}

func (p postgresRegisterRepository) VerifyOTP(cxt context.Context, userId string, pin string) (bool, error) {
	var otp registration.UserOneTimePin
	result := p.conn.Model(registration.UserOneTimePin{
		UserID: userId,
		Pin:    pin,
	}).First(&otp).WithContext(cxt)
	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, errors.New("pin not found")

		} else {
			return false, errors.New("error occurred while fetching pin")

		}
		return false, result.Error
	}
	today := time.Now()
	if otp.ExpiryDate.Before(today) {
		return false, errors.New("otp expired")
	}
	return true, nil

}
func (p postgresRegisterRepository) getUser(ctx context.Context, id string) (*registration.User, error) {
	var user registration.User
	result := p.conn.First(&user, "id = ?", id).WithContext(ctx)
	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("Record not found")

		} else {
			return nil, errors.New("Error occurred while fetching user")

		}
		return nil, result.Error
	}
	return &user, nil
}
func (p postgresRegisterRepository) SaveDetails(cxt context.Context, userId string, details registration.UserDetails) (bool, error) {
	user, error := p.getUser(cxt, userId)
	if error != nil {
		return false, errors.New("user not found")
	}
	err := p.conn.Model(&user).Association(constants.DETAILS).Append(details)
	if err != nil {
		return false, errors.New(err.Error())
	}
	return true, nil
}

func (p postgresRegisterRepository) SaveResidenceDetails(cxt context.Context, userId string, details registration.ResidenceDetails) (bool, error) {
	user, error := p.getUser(cxt, userId)
	if error != nil {
		return false, errors.New("user not found")
	}
	err := p.conn.Model(&user).Association(constants.RESIDENCE).Append(details)
	if err != nil {
		return false, errors.New(err.Error())
	}
	return true, nil
}

func (p postgresRegisterRepository) SaveUserIdentification(cxt context.Context, userId string, identification registration.UserIdentification) (bool, error) {
	user, error := p.getUser(cxt, userId)
	if error != nil {
		return false, errors.New("user not found")
	}
	err := p.conn.Model(&user).Association(constants.IDENTIFICATION).Append(identification)
	if err != nil {
		return false, errors.New(err.Error())
	}
	return true, nil
}

func (p postgresRegisterRepository) SaveUserIncome(cxt context.Context, userId string, income registration.UserIncome) (bool, error) {
	user, error := p.getUser(cxt, userId)
	if error != nil {
		return false, errors.New("user not found")
	}
	err := p.conn.Model(&user).Association(constants.INCOME).Append(income)
	if err != nil {
		return false, errors.New(err.Error())
	}
	return true, nil
}

func (p postgresRegisterRepository) SaveEmployment(cxt context.Context, userId string, employment registration.UserEmployment) (bool, error) {
	user, error := p.getUser(cxt, userId)
	if error != nil {
		return false, errors.New("user not found")
	}
	err := p.conn.Model(&user).Association(constants.EMPLOYMENT).Append(employment)
	if err != nil {
		return false, errors.New(err.Error())
	}
	return true, nil
}

func (p postgresRegisterRepository) SaveContact(cxt context.Context, userId string, contact registration.Contact) (bool, error) {
	user, error := p.getUser(cxt, userId)
	if error != nil {
		return false, errors.New("user not found")
	}
	err := p.conn.Model(&user).Association(constants.CONTACT).Append(contact)
	if err != nil {
		return false, errors.New(err.Error())
	}
	return true, nil
}

func NewPostgresRegisterRepository(conn *gorm.DB) RegisterRepository {
	return &postgresRegisterRepository{
		conn: conn,
	}

}
