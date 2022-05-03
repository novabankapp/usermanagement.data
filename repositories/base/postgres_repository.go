package base

import (
	"context"
	"errors"
	"github.com/novabankapp/common.infrastructure/postgres"
	"gorm.io/gorm"
)

type postgresRepository struct {
	conn *gorm.DB
}

func NewPostGreRepository(conn *gorm.DB) Repository {
	return &postgresRepository{
		conn,
	}
}
func GetById[E Entity](rep *postgresRepository, ctx context.Context, id string) (*E, error) {
	var entity E
	result := rep.conn.First(&entity, "id = ?", id).WithContext(ctx)
	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("Record not found")

		} else {
			return nil, errors.New("Error occurred while fetching entity")

		}
		return nil, result.Error
	}
	return &entity, nil
}
func Create[E Entity](rep *postgresRepository, ctx context.Context, entity E) (*E, error) {

	result := rep.conn.Create(&entity).WithContext(ctx)
	if result.Error != nil && result.RowsAffected != 1 {
		return nil, errors.New("Error occurred while creating a new entity")

	}

	return &entity, nil

}
func Update[E Entity](rep *postgresRepository, ctx context.Context, entity E, id string) (bool, error) {

	// Create a user object
	var value E

	// Read the user which is to be updated
	result := rep.conn.First(&value, "id = ?", id).WithContext(ctx)
	if result.Error != nil {
		return false, errors.New("Error occurred while updating the entity")

	}
	value = entity

	// Save the updated user
	tx := rep.conn.Save(&value)
	if tx.RowsAffected != 1 {
		return false, errors.New("Error occurred while updating entity")
	}
	return true, nil
}
func Delete[E Entity](rep *postgresRepository, ctx context.Context, id string) (bool, error) {
	var value E
	result := rep.conn.First(&value, "id = ?", id)
	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, errors.New("Record not found")
		} else {
			return false, errors.New("Error occurred while deleting entity")
		}
		return false, result.Error
	}

	tx := rep.conn.Delete(&value).WithContext(ctx)
	if tx.RowsAffected != 1 {
		return false, errors.New("Error occurred while deleting entity")

	}
	return true, nil
}
func Get[E Entity](rep *postgresRepository, ctx context.Context,
	page int, pageSize int, query string, orderBy string) (*[]E, error) {

	var values []E
	tx := rep.conn
	if query != "" {
		tx = tx.Where(query)
	}
	if orderBy != "" {
		//Order("created_at asc")
		tx = tx.Order(orderBy)
	}
	tx = tx.Scopes(postgres.Paginate(page, pageSize)).Find(&values).WithContext(ctx)

	if tx.RowsAffected == 0 {
		return nil, errors.New("Read users returned with empty results")
	}
	return &values, nil

}
