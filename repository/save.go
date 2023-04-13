package repository

import (
	"context"
	"fmt"
	"log"
	"time"
)

func (r *Repository) Save(amount float64, from string, to string, rate float64, result float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := r.DB.GetConn(ctx); if err != nil {
		log.Fatalf("Couldn't connect to genesis database: %s", err); return err
	}
	defer conn.Close()
	query := `
		INSERT INTO genesis.exchanges(amount, from_currency, to_currency, rate, conversion_result)	
		VALUES 
			(?,
				(SELECT id FROM currency WHERE currency_name = ? ),
				(SELECT id FROM currency WHERE currency_name = ? ),
			?, ?)`
	
	stmt, err := conn.PrepareContext(ctx, query); if err != nil {
		log.Fatalf("Error Preparing query: %s", err); return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(amount, from, to, rate, result)
	if err != nil {
		return fmt.Errorf("Couldn't insert to genesis database: %s", err)
	}
	return nil
}