package infra

import (
	"context"
	"database/sql"
)

type Datasource struct {
	Source *sql.DB
}

type IDatasource interface {
	Connect(driver string, source string) error
	GetConn(ctx context.Context) (*sql.Conn, error)
	Disconnect() error
}