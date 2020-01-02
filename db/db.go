package db

import "database/sql"
import "fmt"

type CountryDB interface {
	CountAll() (int, error)
}

type countryDB struct {
	db *sql.DB
}

func New(db *sql.DB) *countryDB {
	return &countryDB{
		db: db,
	}
}

// @todo #1 /countries/count: extract SQL query
func (c *countryDB) CountAll() (int, error) {
	var count int

	// There is no check for ErrNoRows because COUNT(*) always returns a single row
	err := c.db.QueryRow("SELECT COUNT(*) FROM countries").Scan(&count)
	if err != nil {
		return count, fmt.Errorf("Scan() has failed: %v", err)
	}
	return count, nil
}
