package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
)

var (
	uninitializedDBError = "ERROR: relation \"meta_info\" does not exist (SQLSTATE 42P01)"
)

func InitDb() error {
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

		_, err = tx.Exec(context.Background(), Schema)
		if err != nil {
			return err
		}

		tx.Commit(context.Background())

	}
	return nil
}

func isInitRequired(conn *pgx.Conn) (bool, error) {
	var version int64
	err := conn.QueryRow(context.Background(), "SELECT * FROM meta_info;").Scan(&version)

	if err != nil {
		if err.Error() == uninitializedDBError {
			return true, nil
		}
		return false, err
	}
	log.Debug("DB VERSION: ", version)
	return false, nil
}
