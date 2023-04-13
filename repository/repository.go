package repository

import (
	"github.com/Obdurat/genesis/infra"
	_ "github.com/go-sql-driver/mysql"
)

func New(db infra.IDatasource) IRepository {
	return &Repository{db}
}

var Repo = New(infra.DB)