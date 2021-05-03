package helpers

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4"
)

type DBHelper struct {
	uri string
}

func NewDBHelper(uri string) DBHelper {
	return DBHelper{uri: uri}
}

func (d DBHelper) Reset() error {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URI"))
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	txn, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}
	defer txn.Rollback(context.Background())

	if err = d.clean(txn); err != nil {
		return err
	}
	if err = d.init(txn); err != nil {
		return err
	}

	return txn.Commit(context.Background())
}

func (d DBHelper) clean(tx pgx.Tx) error {
	_, err := tx.Exec(context.Background(), `
		TRUNCATE TABLE videos CASCADE;
		TRUNCATE TABLE users CASCADE;
	`)
	return err
}

func (d DBHelper) init(tx pgx.Tx) error {
	_, err := tx.Exec(context.Background(), `
		INSERT INTO users (username, hashpass) VALUES ('admin', '$2y$10$XcWmOIgAuT90XB/7cSwK5e1PTEUeJgXcO47Zgjx6RHh2phZVFqc/C');
	`)
	return err
}
