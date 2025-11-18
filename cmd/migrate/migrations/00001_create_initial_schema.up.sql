-- SQL in section 'Up' is executed when this migration is applied
-- 00001_create_initial_schema.up.sql

-- 1. Create Cities Table
CREATE TABLE Cities (
    city_id SERIAL PRIMARY KEY,
    city_name VARCHAR(255) NOT NULL UNIQUE,
    total_budget_co2e DECIMAL(18, 2) NOT NULL, -- The city's total remaining budget in tonnes of CO2e
    budget_set_year INT NOT NULL,
    last_updated TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 2. Create ActivityCategories Table
CREATE TABLE ActivityCategories (
    activity_id SERIAL PRIMARY KEY,
    category_name VARCHAR(100) NOT NULL UNIQUE, -- e.g., 'Transport', 'Buildings', 'Waste'
    description TEXT
);

-- 3. Create AnnualEmissions Table
CREATE TABLE AnnualEmissions (
    emission_id SERIAL PRIMARY KEY,
    city_id INT NOT NULL,
    activity_id INT NOT NULL,
    year INT NOT NULL, -- The calendar year for this emissions data (e.g., 2024)
    annual_emissions_co2e DECIMAL(18, 2) NOT NULL, -- Emissions for this activity in this year (tonnes of CO2e)
    data_source VARCHAR(255),

    -- Define Foreign Keys
    CONSTRAINT fk_city
        FOREIGN KEY (city_id)
        REFERENCES Cities (city_id)
        ON DELETE CASCADE,

    CONSTRAINT fk_activity
        FOREIGN KEY (activity_id)
        REFERENCES ActivityCategories (activity_id)
        ON DELETE RESTRICT,

    -- Ensure we only have one entry per city, per activity, per year
    UNIQUE (city_id, activity_id, year)
);