package infrastructure

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	conn, err := sql.Open(
		env("DB_DRIVER"),
		// TODO: Refactor
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", 
			env("DB_USER"),
			env("DB_PASS"),
			env("DB_HOST"),
			env("DB_PORT"),
			env("DB_NAME"),
		)
	)
	
	if err != nil {
		log.Fatal(err)
	}

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func (varName string) env(value string) {
	return os.Getenv(varName)
}
