package postgres

import (
	"context"
	"errors"
	domainbase "github.com/novabankapp/usermanagement.data/domain/base"

	"github.com/novabankapp/common.infrastructure/postgres"
	"github.com/novabankapp/usermanagement.data/repositories/base"
	"gorm.io/gorm"
)

type postgresRepository[E domainbase.Entity] struct {
	conn *gorm.DB
}

func NewPostGreRepository(conn *gorm.DB) base.Repository {
	return &postgresRepository{
		conn,
	}
}
func (rep *postgresRepository[E]) GetById(ctx context.Context, id string) (*E, error) {
	var entity E
	result := rep.conn.First(&entity, "id = ?", id).WithContext(ctx)
	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("Record not found")

		} else {
			return nil, errors.New("Error occurred while fetching entity")

		}

	}
	return &entity, nil
}
func (rep *postgresRepository[E]) Create(ctx context.Context, entity E) (*E, error) {

	result := rep.conn.Create(&entity).WithContext(ctx)
	if result.Error != nil && result.RowsAffected != 1 {
		return nil, errors.New("Error occurred while creating a new entity")

	}

	return &entity, nil

}
func (rep *postgresRepository[E]) Update(ctx context.Context, entity E, id string) (bool, error) {

	// Create a user object
	var value E

	// Read the user which is to be updated
	result := rep.conn.First(&value, "id = ?", id).WithContext(ctx)
	if result.Error != nil {
		return false, errors.New("error occurred while updating the entity")

	}
	value = entity

	// Save the updated user
	tx := rep.conn.Save(&value)
	if tx.RowsAffected != 1 {
		return false, errors.New("error occurred while updating entity")
	}
	return true, nil
}
func (rep *postgresRepository[E]) Delete(ctx context.Context, id string) (bool, error) {
	var value E
	result := rep.conn.First(&value, "id = ?", id)
	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, errors.New("record not found")
		} else {
			return false, errors.New("error occurred while deleting entity")
		}

	}

	tx := rep.conn.Delete(&value).WithContext(ctx)
	if tx.RowsAffected != 1 {
		return false, errors.New("error occurred while deleting entity")

	}
	return true, nil
}
func (rep *postgresRepository[E]) Get(ctx context.Context,
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