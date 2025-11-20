package types

// BudgetYear represents the budget state for a single year
type BudgetYear struct {
	Year              int     `json:"year"`
	AnnualEmissions   float64 `json:"annual_emissions_co2e"`
	RemainingBudget   float64 `json:"remaining_budget_co2e"`
	BudgetConsumedPct float64 `json:"budget_consumed_percent"`
}

// BudgetProjection holds the full budget consumption projection for a city
type BudgetProjection struct {
	CityID              int          `json:"city_id"`
	CityName            string       `json:"city_name"`
	InitialBudget       float64      `json:"initial_budget_co2e"`
	BudgetSetYear       int          `json:"budget_set_year"`
	YearlyBreakdown     []BudgetYear `json:"yearly_breakdown"`
	YearsUntilExhausted int          `json:"years_until_budget_exhausted"`
}
