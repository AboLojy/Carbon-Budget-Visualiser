package budgetcalculator

import (
	"context"
	"database/sql"
	"time"

	"github.com/AboLojy/Carbon-Budget-Visualiser/types"
)

// Store defines persistence operations for budget calculations.
type BudgetStore interface {
	CalculateBudgetConsumption(ctx context.Context, cityID, year int) (*types.BudgetProjection, error)
	CreateCity(ctx context.Context, c *types.City) (*types.City, error)
}

// Store implements Store using *sql.DB.
type Store struct {
	db *sql.DB
}

// NewStore creates a Store backed by the provided DB.
func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// CalculateBudgetConsumption fetches city data and annual emissions, then calculates year-by-year budget remaining.
func (s *Store) CalculateBudgetConsumption(ctx context.Context, cityID, year int) (*types.BudgetProjection, error) {
	// Fetch city info
	cityQuery := `SELECT city_id, city_name, total_budget_co2e, budget_set_year FROM Cities WHERE city_id = $1`
	row := s.db.QueryRowContext(ctx, cityQuery, cityID)

	var city types.City
	if err := row.Scan(&city.CityID, &city.CityName, &city.TotalBudgetCO2E, &city.BudgetSetYear); err != nil {
		return nil, err
	}

	// Fetch annual emissions for this city, ordered by year
	emissionsQuery := `SELECT year, SUM(annual_emissions_co2e) as total_emissions 
	                   FROM AnnualEmissions 
	                   WHERE city_id = $1 AND year = $2
	                   GROUP BY year
					   limit(1)`

	rows, err := s.db.QueryContext(ctx, emissionsQuery, cityID, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projection := &types.BudgetProjection{
		CityID:          city.CityID,
		CityName:        city.CityName,
		InitialBudget:   city.TotalBudgetCO2E,
		BudgetSetYear:   city.BudgetSetYear,
		YearlyBreakdown: []types.BudgetYear{},
	}

	for rows.Next() {
		var year int
		var annualEmissions float64
		if err := rows.Scan(&year, &annualEmissions); err != nil {
			return nil, err
		}

		projection.YearlyBreakdown = calculateBudgetConsumption(year, city.TotalBudgetCO2E, annualEmissions)
		projection.YearsUntilExhausted = len(projection.YearlyBreakdown)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return projection, nil
}
func calculateBudgetConsumption(year int, totalBudgetCO2E, annualEmissions float64) []types.BudgetYear {

	yearlyBreakdown := []types.BudgetYear{}

	remainingBudget := totalBudgetCO2E
	yearsUntilExhausted := year - 1

	for remainingBudget > 0 {

		yearsUntilExhausted++
		remainingBudget -= annualEmissions
		if remainingBudget < 0 {
			remainingBudget = 0
		}
		consumedPct := ((totalBudgetCO2E - remainingBudget) / totalBudgetCO2E) * 100

		yearData := types.BudgetYear{
			Year:              yearsUntilExhausted,
			AnnualEmissions:   annualEmissions,
			RemainingBudget:   remainingBudget,
			BudgetConsumedPct: consumedPct,
		}
		yearlyBreakdown = append(yearlyBreakdown, yearData)
	}
	return yearlyBreakdown
}

func (s *Store) CreateCity(ctx context.Context, c *types.City) (*types.City, error) {
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
