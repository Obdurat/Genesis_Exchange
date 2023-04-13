package repository

import "github.com/Obdurat/genesis/infra"

type Exchange struct {
	Id int `json:"id"`
	Amount float64 `json:"amount"`
	From string `json:"from"`
	To string `json:"to"`
	Rate float64 `json:"rate"`
	Result float64 `json:"result"`
}

type Repository struct {
	DB infra.IDatasource
}

type IRepository interface {
	Save(amount float64, from string, to string, rate float64, result float64) error
	List(offset int64) ([]Exchange, error)
}