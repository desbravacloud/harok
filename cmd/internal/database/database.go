package database

import (
	"database/sql"
	"fmt"

	"github.com/Spotlitebr/rocket/cmd/internal/driver"
)

var DB *sql.DB

var dsn = "host=localhost port=5432 dbname=test user=test password=test"

func InsertIntoAppTable(app App) {

	DB, err := driver.InitializeConnectionSQL(dsn)
	if err != nil {
		println(err)
	}
	defer DB.SQL.Close()
	_, err = DB.SQL.Exec("INSERT INTO apps (id, name, hostname, language, coderepo, imagerepo, created_at, updated_at) values (DEFAULT, $1, $2, $3, $4, $5, now(), now())", app.Name, app.Hostname, app.Language, app.CodeRepo, app.ImageRepo)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("App registered successfully!")
}
