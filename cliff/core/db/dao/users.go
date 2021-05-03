package dao

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type UserDAO struct {
	conn *pgxpool.Pool
}

func NewUserDAO(conn *pgxpool.Pool) *UserDAO {
	return &UserDAO{conn}
}

func (u *UserDAO) Create(username string) (int, error) {
	commandTag, err := u.conn.Exec(context.Background(), "INSERT INTO users (username) VALUES ($1);", username)
	if err != nil {
		return 0, err
	}
	if commandTag.RowsAffected() != 1 {
		return 0, errors.New("failed to create new user")
	}
	return u.Login(username)
}

func (u *UserDAO) Login(username string) (int, error) {
	var id int
	err := u.conn.QueryRow(context.Background(), "SELECT id FROM users WHERE username=$1;", username).Scan(&id)
	return id, err
}

func (u *UserDAO) Get(username string) (*User, error) {
	var user User
	err := u.conn.QueryRow(context.Background(), "SELECT id, username FROM users WHERE username=$1;", username).Scan(&user.ID, &user.Username)
	return &user, err
}
