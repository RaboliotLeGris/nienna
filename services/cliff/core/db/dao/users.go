package dao

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
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

func (u *UserDAO) Create(username, password string) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return 0, fmt.Errorf("fail to hash the password")
	}

	commandTag, err := u.conn.Exec(context.Background(), "INSERT INTO users (username, hashpass) VALUES ($1, $2);", username, string(hashedPassword))
	if err != nil {
		return 0, err
	}
	if commandTag.RowsAffected() != 1 {
		return 0, errors.New("failed to create new user")
	}
	return u.Login(username, password)
}

func (u *UserDAO) Login(username, password string) (int, error) {
	var id int
	var hashpass string
	if err := u.conn.QueryRow(context.Background(), "SELECT id, hashpass FROM users WHERE username=$1;", username).Scan(&id, &hashpass); err != nil {
		return 0, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashpass), []byte(password)); err != nil {
		return 0, err
	}

	return id, nil
}

func (u *UserDAO) Get(username string) (*User, error) {
	var user User
	err := u.conn.QueryRow(context.Background(), "SELECT id, username FROM users WHERE username=$1;", username).Scan(&user.ID, &user.Username)
	return &user, err
}
