package postgres

import (
	"database/sql"
	"gowithpg/config"
	"gowithpg/internal/model"

	_ "github.com/lib/pq"
)

type Postgres struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Postgres, error) {
	db, err := sql.Open("postgres", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS students (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT UNIQUE,
		age INTEGER
	)
	`)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		Db: db,
	}, nil
}
// CREATE
func (p *Postgres) CreateStock(stock *model.Stock) error {
	query := `
	INSERT INTO stocks (name, symbol, price, quantity)
	VALUES ($1, $2, $3, $4)
	RETURNING id
	`

	return p.Db.QueryRow(
		query,
		stock.Name,
		stock.Symbol,
		stock.Price,
		stock.Quantity,
	).Scan(&stock.ID)
}

// GET ONE
func (p *Postgres) GetStock(id int) (*model.Stock, error) {
	var stock model.Stock

	query := `
	SELECT id, name, symbol, price, quantity
	FROM stocks
	WHERE id=$1
	`

	err := p.Db.QueryRow(query, id).Scan(
		&stock.ID,
		&stock.Name,
		&stock.Symbol,
		&stock.Price,
		&stock.Quantity,
	)

	if err != nil {
		return nil, err
	}

	return &stock, nil
}

// GET ALL
func (p *Postgres) GetAllStocks() ([]model.Stock, error) {
	query := `
	SELECT id, name, symbol, price, quantity
	FROM stocks
	`

	rows, err := p.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stocks []model.Stock

	for rows.Next() {
		var stock model.Stock

		err := rows.Scan(
			&stock.ID,
			&stock.Name,
			&stock.Symbol,
			&stock.Price,
			&stock.Quantity,
		)

		if err != nil {
			return nil, err
		}

		stocks = append(stocks, stock)
	}

	return stocks, nil
}


// UPDATE
func (p *Postgres) UpdateStock(id int, stock *model.Stock) error {
	query := `
	UPDATE stocks
	SET name=$1, symbol=$2, price=$3, quantity=$4
	WHERE id=$5
	`

	_, err := p.Db.Exec(
		query,
		stock.Name,
		stock.Symbol,
		stock.Price,
		stock.Quantity,
		id,
	)

	return err
}

// DELETE
func (p *Postgres) DeleteStock(id int) error {
	query := `DELETE FROM stocks WHERE id=$1`

	_, err := p.Db.Exec(query, id)

	return err
}