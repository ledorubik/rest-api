package postgres

import (
	"context"
	"log"
	"rest-api/pkg/models"

	"gorm.io/gorm"
)

type User struct {
	ID    uint
	Uuid  string
	Login string
	gorm.Model
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r UserRepository) CreateUser(ctx context.Context, user *models.User) (*User, error) {
	model := toPostgresUser(user)
	err := r.db.WithContext(ctx).Create(model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (r UserRepository) GetUserByUuid(ctx context.Context, uuid string) (*models.User, error) {
	user := new(User)
	err := r.db.WithContext(ctx).Where(&User{
		Uuid: uuid,
	}).First(&user).Error

	if err != nil {
		return nil, err
	}

	return toModel(user), nil
}

func (r UserRepository) GetUserByLogin(ctx context.Context, login string) (*models.User, error) {
	user := new(User)
	err := r.db.WithContext(ctx).Where(&User{
		Login: login,
	}).First(&user).Error

	if err != nil {
		return nil, err
	}

	return toModel(user), nil
}

func (r UserRepository) UpdateUsersLogin(ctx context.Context, user *models.User) (*User, error) {
	model := toPostgresUser(user)
	err := r.db.WithContext(ctx).Model(&User{}).Where(&User{
		Uuid: user.Uuid,
	}).Update("login", user.Login).First(&user).Error

	if err != nil {
		log.Printf("UPDATE_USER_ERROR: %v", err)
		return nil, err
	}
	return model, nil
}

func (r UserRepository) DeleteUser(ctx context.Context, user *models.User) error {
	err := r.db.WithContext(ctx).Where(&User{
		Uuid: user.Uuid,
	}).Delete(&user)

	if err != nil {
		log.Printf("DELETE_USER_ERROR: %v", err)
		return err.Error
	}

	return nil
}

func toPostgresUser(user *models.User) *User {
	return &User{
		Login: user.Login,
		Uuid:  user.Uuid,
	}
}

func toModel(user *User) *models.User {
	return &models.User{
		ID:    user.ID,
		Login: user.Login,
		Uuid:  user.Uuid,
		Model: gorm.Model{
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		},
	}
}
