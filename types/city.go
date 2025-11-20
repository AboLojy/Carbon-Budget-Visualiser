package types

import "time"

// City represents a city record for the API and database
type City struct {
    CityID          int       `json:"city_id,omitempty" db:"city_id"`
    CityName        string    `json:"city_name" db:"city_name" validate:"required"`
    TotalBudgetCO2E float64   `json:"total_budget_co2e" db:"total_budget_co2e" validate:"required,gt=0"`
    BudgetSetYear   int       `json:"budget_set_year" db:"budget_set_year" validate:"required,gte=1900,lte=2100"`
    LastUpdated     time.Time `json:"last_updated,omitempty" db:"last_updated"`
}
