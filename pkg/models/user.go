package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID         uint
	Uuid       string
	Login      string
	gorm.Model //@todo
}

//type UserRepository interface {
//	CreateUser(ctx context.Context, user *User) error
//	GetUser(ctx context.Context, login string) (*User, error)
//}
