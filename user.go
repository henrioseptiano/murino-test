package main

import (
	"context"
	"time"
)

type UserID int64

type User struct {
	ID         UserID    `json:"id"`
	RoleID     int64     `json:"role_id"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Name       string    `json:"name"`
	LastAccess time.Time `json:"last_access"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

type UserRepository interface {
	CheckUser(context.Context, string) (*User, error)
	//CreateUser(context.Context, User) error
	//UpdateUser(context.Context, User, UserID) error
	//DeleteUser(context.Context, UserID) error
	//GetAllUser(context.Context) ([]*User, error)
}
