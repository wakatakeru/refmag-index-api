package database

type SqlHandler interface {
	Execute(string, ...interface{}) (Result, error)
	Execute(string, ...interface{}) (Row, error)
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Row interface {
	Scan(...interfance{}) error
	Next() bool
	Close() error
}
