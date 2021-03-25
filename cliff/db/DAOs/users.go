package daos

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
)

type User struct {
	ID       int
	Username string
}

type UserDAO struct {
	conn *pgxpool.Pool
}

func NewUserDAO(conn *pgxpool.Pool) *UserDAO {
	return &UserDAO{conn}
}

func (u *UserDAO) Create(username string) error {
	commandTag, err := u.conn.Exec(context.Background(), "INSERT INTO users (username) VALUES ($1);", username)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return errors.New("failed to create new user")
	}
	return nil
}

func (u *UserDAO) Login(username string) error {
	var _username string
	return u.conn.QueryRow(context.Background(), "SELECT username FROM users WHERE username=$1;", username).Scan(&_username)
}

func (u *UserDAO) Get(username string) (*User, error) {
	var user User
	err := u.conn.QueryRow(context.Background(), "SELECT id, username FROM users WHERE username=$1;", username).Scan(&user.ID, &user.Username)
	return &user, err
}
