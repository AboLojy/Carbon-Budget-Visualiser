package emission

import (
	"context"
	"database/sql"
	"time"

	"github.com/AboLojy/Carbon-Budget-Visualiser/types"
)

// Store defines persistence operations for emission-related entities.
type CityEmissionStore interface {
	CreateCityEmission(ctx context.Context, c *types.City) (*types.City, error)
}

type Store struct {
	db *sql.DB
}

// NewCityEmissionStore returns a new CityEmissionStore.
func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// CreateCity inserts a city record and returns the created entity.
func (s *Store) CreateCityEmission(ctx context.Context, c *types.City) (*types.City, error) {
	query := `INSERT INTO Cities (city_name, total_budget_co2e, budget_set_year) VALUES ($1, $2, $3) RETURNING city_id, last_updated`

	row := s.db.QueryRowContext(ctx, query, c.CityName, c.TotalBudgetCO2E, c.BudgetSetYear)
	var id int
	var lastUpdated time.Time
	if err := row.Scan(&id, &lastUpdated); err != nil {
		return nil, err
	}
	c.CityID = id
	c.LastUpdated = lastUpdated
	return c, nil
}
