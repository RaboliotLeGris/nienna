package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func InitDB() error {
	log.Info("DB - init - Checking database status")
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URI"))
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	needInit, err := isInitRequired(conn)
	if err != nil {
		return err
	}

	if needInit {
		log.Info("DB - init - Initializing database")
		tx, err := conn.Begin(context.Background())
		if err != nil {
			return err
		}

		// Create the schema
		if _, err = tx.Exec(context.Background(), Schema); err != nil {
			return err
		}

		// Set default value in DB
		var adminPassword string
		if os.Getenv("NIENNA_ADMIN_PASSWORD") != "" {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("NIENNA_ADMIN_PASSWORD")), 10)
			if err != nil {
				return fmt.Errorf("fail to hash the password")
			}
			adminPassword = string(hashedPassword)
		} else if os.Getenv("NIENNA_DEV") == "true" {
			adminPassword = "$2y$10$XcWmOIgAuT90XB/7cSwK5e1PTEUeJgXcO47Zgjx6RHh2phZVFqc/C"
		} else {
			return fmt.Errorf("no admin password provided (env: %s)", os.Getenv("NIENNA_DEV"))
		}

		if _, err = tx.Exec(context.Background(), "INSERT INTO users (username, hashpass) VALUES ('admin', $1);", adminPassword); err != nil {
			return err
		}

		if err = tx.Commit(context.Background()); err != nil {
			return err
		}

	}
	return nil
}

func isInitRequired(conn *pgx.Conn) (bool, error) {
	var version int64
	if err := conn.QueryRow(context.Background(), "SELECT * FROM meta_info;").Scan(&version); err != nil {
		if err.Error() == "ERROR: relation \"meta_info\" does not exist (SQLSTATE 42P01)" {
			return true, nil
		}
		return false, err
	}
	log.Debug("DB VERSION: ", version)
	return false, nil
}
