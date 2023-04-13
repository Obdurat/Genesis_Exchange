package infra

import (
	"context"
	"database/sql"
)

func (d *Datasource) Connect(driver string, source string) error {
	db, err := sql.Open(driver, source); if err != nil {
		return err
	}
	err = db.Ping(); if err != nil {
		return err
	}
	d.Source = db
	return nil
}

func (d *Datasource) Disconnect() error {
	if err := d.Source.Close(); err != nil {
		return err
	}
	return nil
}

func (d *Datasource) GetConn(ctx context.Context) (*sql.Conn, error) {
	return d.Source.Conn(ctx)
}

func New() IDatasource {
	return &Datasource{}
}

var DB = New()