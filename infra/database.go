package infra

import "database/sql"

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

func (d *Datasource) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := d.Source.Query(query, args...); if err != nil {
		return nil, err
	}
	return rows, nil
}

func (d *Datasource) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := d.Source.Exec(query, args...); if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *Datasource) Disconnect() error {
	if err := d.Source.Close(); err != nil {
		return err
	}
	return nil
}

func New() IDatasource {
	return &Datasource{}
}

var DB = New()