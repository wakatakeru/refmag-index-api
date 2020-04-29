package infrastructure

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wakatakeru/refmag-index-api/interfaces/database"
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
		log.Panic(err)
	}

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err 
	}
	res.Result = result
	return res, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()	
}

func (r SqlHandler) RowAffected() (int64, error) {
	return r.Result.RowAffected()
}

type SqlRow struct {
	Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Row.Scan(dest...)
}

func (r SqlRow) Next() bool {
    return r.Rows.Next()
}

func (r SqlRow) Close() error {
    return r.Rows.Close()
}

func env(varName string) string {
	return os.Getenv(varName)
}
