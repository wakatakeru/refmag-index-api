package database

type SqlHandler interface {
	Execute(string, ...interface{}) (SqlResult, error)
	Query(string, ...interface{}) (Row, error)
}

type SqlResult interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}
