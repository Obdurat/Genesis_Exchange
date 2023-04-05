package infra

import "database/sql"

type Datasource struct {
	Source *sql.DB
}

type IDatasource interface {
	Connect(driver string, source string) error
	Query(query string, args interface{}) (*sql.Rows, error)
	Exec(query string, args interface{}) (sql.Result, error) 
}

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

func (d *Datasource) Query(query string, args interface{}) (*sql.Rows, error) {
	rows, err := d.Source.Query(query, args); if err != nil {
		return nil, err
	}
	return rows, nil
}

func (d *Datasource) Exec(query string, args interface{}) (sql.Result, error) {
	result, err := d.Source.Exec(query, args); if err != nil {
		return nil, err
	}
	return result, nil
}