package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/Obdurat/genesis/infra"
)

type Exchange struct {
	Id *int `json:"id"`
	Amount *float64 `json:"amount"`
	From *string `json:"from"`
	To *string `json:"to"`
	Rate *float64 `json:"rate"`
	Result *float64 `json:"result"`
}

func Save(amount float64, from string, to string, rate float64, result float64) {
	db := infra.DB
	if err := db.Connect("mysql", fmt.Sprintf("root:%s@tcp(db:3306)/genesis", os.Getenv("MYSQL_ROOT_PASSWORD"))); err != nil {
		log.Fatalf("Couldn't connect to genesis database: %s", err); return
	}	

	defer db.Disconnect()

	query := `
		INSERT INTO genesis.exchanges(amount, from_currency, to_currency, rate, conversion_result)	
		VALUES 
			(?,
				(SELECT id FROM currency WHERE currency_name = ? ),
				(SELECT id FROM currency WHERE currency_name = ? ),
			?, ?)`

	_, err := db.Exec(query, amount, from, to, rate, result)
	if err != nil {
		log.Fatalf("Couldn't insert to genesis database: %s", err); return
	}
}

func List() ([]Exchange, error) {
	db := infra.DB
	if err := db.Connect("mysql", fmt.Sprintf("root:%s@tcp(db:3306)/genesis", os.Getenv("MYSQL_ROOT_PASSWORD"))); err != nil {
		log.Fatalf("Couldn't connect to genesis database: %s", err); return nil, err
	}	

	defer db.Disconnect()
	query := `SELECT 
		exchanges.id,
		exchanges.amount,
		from_currency.currency_name as "from",
		to_currency.currency_name as "to",
		exchanges.rate,
		exchanges.conversion_result
		FROM exchanges
		INNER JOIN currency as from_currency ON exchanges.from_currency = from_currency.id
		INNER JOIN currency as to_currency ON exchanges.to_currency = to_currency.id;`

	var out []Exchange

	results, err := db.Query(query)

	for results.Next() {
		var exchange Exchange
		results.Scan(&exchange.Id, &exchange.Amount, &exchange.From, &exchange.To, &exchange.Rate, &exchange.Result)
		out = append(out, exchange)
	}

	if err != nil {
		log.Fatalf("Couldn't insert to genesis database: %s", err); return nil, err
	}

	return out, nil
}