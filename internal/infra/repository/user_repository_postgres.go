package repository

import (
	"github.com/henriquemarlon/pond-micro-cam/backend/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

type UserRepositoryPostgres struct {
	db *sqlx.DB
}

func NewUserRepositoryPostgres(db *sqlx.DB) *UserRepositoryPostgres {
	return &UserRepositoryPostgres{db: db}
}

// User methods
func (r *UserRepositoryPostgres) CreateUser(user *entity.User) (*entity.User, error) {
	var userCreated entity.User
	err := r.db.QueryRow("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id, name, email, created_at", user.Name, user.Email, user.Password).Scan(&userCreated.ID, &userCreated.Name, &userCreated.Email, &userCreated.CreatedAt)
	if err != nil {
		return nil, err
	}

	// Log user creation
	if err := r.logAction(userCreated.ID, "create_user"); err != nil {
		return nil, err
	}

	return &userCreated, nil
}

func (r *UserRepositoryPostgres) FindAllUsers() ([]*entity.User, error) {
	var users []*entity.User
	err := r.db.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepositoryPostgres) FindUserById(id string) (*entity.User, error) {
	var user entity.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryPostgres) FindUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}

	// Log action
	if err := r.logAction(user.ID, "user_login"); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryPostgres) logAction(userID, action string) error {
	_, err := r.db.Exec("INSERT INTO logs (user_id, action) VALUES ($1, $2)", userID, action)
	return err
}
