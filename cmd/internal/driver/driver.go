package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

func InitializeConnectionSQL(dsn string) (*DB, error) {
	d, err := openConn(dsn)
	if err != nil {
		fmt.Println("Could not connect to database")
		panic(err)
	}
	dbConn.SQL = d
	return dbConn, nil
}

func openConn(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
