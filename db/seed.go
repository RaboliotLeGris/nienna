package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	if os.Getenv("NIENNA_DEV") == "true" {
		log.SetLevel(log.DebugLevel)
	}

	config, err := parseConfig()
	if err != nil {
		log.Fatal(err)
	}
	if err := InitDB(config); err != nil {
		log.Fatal(err)
	}
}

func InitDB(config Config) error {
	log.Info("DB - initDB - Checking database status")
	conn, err := pgx.Connect(context.Background(), config.db_uri)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	needInit, err := isInitRequired(conn)
	if err != nil {
		return err
	}

	if needInit {
		log.Info("DB - initDB - Initializing database")
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
		if config.admin_pwd != "" {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(config.admin_pwd), 10)
			if err != nil {
				return fmt.Errorf("fail to hash the password")
			}
			adminPassword = string(hashedPassword)
		} else if config.dev_mod {
			adminPassword = "$2y$10$XcWmOIgAuT90XB/7cSwK5e1PTEUeJgXcO47Zgjx6RHh2phZVFqc/C"
		} else {
			return fmt.Errorf("no admin password provided (env: %v)", config.dev_mod)
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
