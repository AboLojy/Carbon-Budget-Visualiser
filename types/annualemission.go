package types

// AnnualEmission represents an annual emissions record
type AnnualEmission struct {
	EmissionID          int     `json:"emission_id,omitempty" db:"emission_id"`
	CityID              int     `json:"city_id" db:"city_id" validate:"required"`
	ActivityID          int     `json:"activity_id" db:"activity_id" validate:"required"`
	Year                int     `json:"year" db:"year" validate:"required,gte=1900,lte=2100"`
	AnnualEmissionsCO2E float64 `json:"annual_emissions_co2e" db:"annual_emissions_co2e" validate:"required,gte=0"`
	DataSource          string  `json:"data_source,omitempty" db:"data_source"`
}
