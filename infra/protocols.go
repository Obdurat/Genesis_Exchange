package infra

import "database/sql"

type Datasource struct {
	Source *sql.DB
}

type IDatasource interface {
	Connect(driver string, source string) error
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Disconnect() error
}