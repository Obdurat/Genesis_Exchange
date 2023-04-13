package repository

import (
	"context"
	"fmt"
	"log"
	"time"
)

func (r *Repository) List(offset int64) ([]Exchange, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := r.DB.GetConn(ctx); if err != nil {
		log.Fatalf("Couldn't connect to genesis database: %s", err); return nil, err
	}
	defer conn.Close()
	
	calcOffset := (offset - 1) * 100
	
	query := fmt.Sprintf(`SELECT 
	exchanges.id,
	exchanges.amount,
	from_currency.currency_name as "from",
	to_currency.currency_name as "to",
	exchanges.rate,
	exchanges.conversion_result
	FROM exchanges
	INNER JOIN currency as from_currency ON exchanges.from_currency = from_currency.id
	INNER JOIN currency as to_currency ON exchanges.to_currency = to_currency.id
	LIMIT 100 OFFSET %d`, calcOffset)
	
	stmt, err := conn.PrepareContext(ctx, query); if err != nil {
		log.Fatalf("Error Preparing query: %s", err); return nil, err
	}

	defer stmt.Close()

	results, err := stmt.Query(); if err != nil {
		log.Fatalf("Error running query: %s", err); return nil, err
	}

	out := make([]Exchange, 0, 100)

	for results.Next() {
		var exchange Exchange
		results.Scan(&exchange.Id, &exchange.Amount, &exchange.From, &exchange.To, &exchange.Rate, &exchange.Result)
		out = append(out, exchange)
	}

	results.Close()

	if err != nil {
		log.Fatalf("Couldn't insert to genesis database: %s", err); return nil, err
	}
	return out, nil
}