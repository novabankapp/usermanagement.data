package users

import (
	"context"
	"errors"
	"fmt"
	"github.com/novabankapp/common.infrastructure/postgres"
	"github.com/novabankapp/usermanagement.data/domain/registration"
	"gorm.io/gorm"
)

type postgresUserRepository struct {
	conn *gorm.DB
}

func NewPostGreRepository(conn *gorm.DB) UserRepository {
	return &postgresUserRepository{
		conn,
	}
}
func (rep *postgresUserRepository) GetUser(ctx context.Context, id string) (*registration.User, error) {
	var user registration.User
	result := rep.conn.First(&user, "id = ?", id).WithContext(ctx)
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
func (rep *postgresUserRepository) Create(ctx context.Context, user registration.User) (*string, error) {

	user.FillDefaults()
	result := rep.conn.Create(&user).WithContext(ctx)
	if result.Error != nil && result.RowsAffected != 1 {
		return nil, errors.New("error occurred while creating a new user")

	}

	return &user.ID, nil

}
func (rep *postgresUserRepository) Update(ctx context.Context, user registration.User) (bool, error) {

	//user.FillDefaults()

	// Create a user object
	var value registration.User

	// Read the user which is to be updated
	result := rep.conn.First(&value, "id = ?", user.ID).WithContext(ctx)
	if result.Error != nil {
		return false, errors.New("Error occurred while updating the user")

	}
	// Update the desired values using the request payload
	value = user

	// Save the updated user
	tx := rep.conn.Save(&value)
	if tx.RowsAffected != 1 {
		return false, errors.New("error occurred while updating user")
	}
	return true, nil
}
func (rep *postgresUserRepository) Delete(ctx context.Context, user registration.User) (bool, error) {
	userId := registration.UserID{ID: user.ID}
	result := rep.conn.First(&user, "id = ?", userId.ID)
	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, errors.New("Record not found")
		} else {
			return false, errors.New("Error occurred while deleting user")
		}
		return false, result.Error
	}

	tx := rep.conn.Delete(&user).WithContext(ctx)
	if tx.RowsAffected != 1 {
		return false, errors.New("Error occurred while deleting user")

	}
	return true, nil
}
func (rep *postgresUserRepository) GetUsers(ctx context.Context, page int, pageSize int, query *string, orderBy *string) (*[]registration.User, error) {
	var users []registration.User

	tx := rep.conn

	if query != nil {
		//Where("name = 'jinni'")
		tx = tx.Where(query)
	}
	if orderBy != nil {
		//Order("created_at asc")
		tx = tx.Order(orderBy)
	}
	tx = tx.Scopes(postgres.Paginate(page, pageSize)).Find(&users).WithContext(ctx)
	fmt.Println(tx.RowsAffected)
	if tx.RowsAffected == 0 {
		return nil, errors.New("read users returned with empty results")
	}
	return &users, nil

}
