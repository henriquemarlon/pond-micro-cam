package entity

import (
	"time"
)

type UserRepository interface {
	CreateUser(user *User) (*User, error)
	FindUserById(id string) (*User, error)
	FindAllUsers() ([]*User, error)
	FindUserByEmail(email string) (*User, error)
}

// type Role string

// const (
// 	AdminRole     Role = "admin"
// 	UserRole      Role = "user"
// 	CollectorRole Role = "collector"
// 	ManagerRole   Role = "manager"
// )

type User struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewUser(name string, email string, password string) *User {
	return &User{
		Name:      name,
		Email:     email,
		Password:  password,
	}
}
