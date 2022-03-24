package repositories

import (
	"context"
	"errors"
	"github.com/novabankapp/usermanagement/usermanagement.data/database/postgres"
	"github.com/novabankapp/usermanagement/usermanagement.data/domain"
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
func (rep *postgresUserRepository) GetUser(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
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
func (rep *postgresUserRepository) Create(ctx context.Context, user domain.User) (*string, error) {

	user.FillDefaults()
	result := rep.conn.Create(&user).WithContext(ctx)
	if result.Error != nil && result.RowsAffected != 1 {
		return nil, errors.New("Error occurred while creating a new user")

	}

	return &user.ID, nil

}
func (rep *postgresUserRepository) Update(ctx context.Context, user domain.User) (bool, error) {

	user.FillDefaults()

	// Create a user object
	var value domain.User

	// Read the user which is to be updated
	result := rep.conn.First(&value, "id = ?", user.ID).WithContext(ctx)
	if result.Error != nil {
		return false, errors.New("Error occurred while updating the user")

	}
	// Update the desired values using the request payload
	value.FirstName = user.FirstName
	value.LastName = user.LastName

	// Save the updated user
	tx := rep.conn.Save(&value)
	if tx.RowsAffected != 1 {
		return false, errors.New("Error occurred while updating user")
	}
	return true, nil
}
func (rep *postgresUserRepository) Delete(ctx context.Context, user domain.User) (bool, error) {
	var userId domain.UserID
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
func (rep *postgresUserRepository) GetUsers(ctx context.Context, page int, pageSize int, query string, orderBy string) (*[]domain.User, error) {
	var users []domain.User

	tx := rep.conn
	if query != "" {
		tx = tx.Where(query)
	}
	if orderBy != "" {
		//Order("created_at asc")
		tx = tx.Order(orderBy)
	}
	tx = tx.Scopes(postgres.Paginate(page, pageSize)).Find(&users).WithContext(ctx)

	if tx.RowsAffected == 0 {
		return nil, errors.New("Read users returned with empty results")
	}
	return &users, nil

}
