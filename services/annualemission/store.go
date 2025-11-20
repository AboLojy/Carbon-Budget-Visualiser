package annualemission

import (
	"context"
	"database/sql"

	"github.com/AboLojy/Carbon-Budget-Visualiser/types"
)

// Store defines persistence operations for annual emissions.
type AnnualEmissionStore interface {
	CreateAnnualEmission(ctx context.Context, ae *types.AnnualEmission) (*types.AnnualEmission, error)
	GetAll(ctx context.Context) ([]types.AnnualEmission, error)
	DeleteByID(ctx context.Context, emissionID int) error
}

// Store implements Store using *sql.DB.
type Store struct {
	db *sql.DB
}

// NewStore creates a Store backed by the provided DB.
func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// CreateAnnualEmission inserts a new annual emission record and returns it with ID.
func (s *Store) CreateAnnualEmission(ctx context.Context, ae *types.AnnualEmission) (*types.AnnualEmission, error) {
	const query = `INSERT INTO AnnualEmissions (city_id, activity_id, year, annual_emissions_co2e, data_source) 
	              VALUES ($1, $2, $3, $4, $5) 
	              RETURNING emission_id`

	row := s.db.QueryRowContext(ctx, query, ae.CityID, ae.ActivityID, ae.Year, ae.AnnualEmissionsCO2E, ae.DataSource)
	if err := row.Scan(&ae.EmissionID); err != nil {
		return nil, err
	}
	return ae, nil
}

// GetAll returns all annual emission records.
func (s *Store) GetAll(ctx context.Context) ([]types.AnnualEmission, error) {
	const query = `SELECT emission_id, city_id, activity_id, year, annual_emissions_co2e, data_source 
	              FROM AnnualEmissions 
	              ORDER BY year DESC, city_id, activity_id`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []types.AnnualEmission
	for rows.Next() {
		var ae types.AnnualEmission
		if err := rows.Scan(&ae.EmissionID, &ae.CityID, &ae.ActivityID, &ae.Year, &ae.AnnualEmissionsCO2E, &ae.DataSource); err != nil {
			return nil, err
		}
		out = append(out, ae)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteByID removes an annual emission record by ID.
func (s *Store) DeleteByID(ctx context.Context, emissionID int) error {
	const query = `DELETE FROM AnnualEmissions WHERE emission_id = $1`
	_, err := s.db.ExecContext(ctx, query, emissionID)
	return err
}
