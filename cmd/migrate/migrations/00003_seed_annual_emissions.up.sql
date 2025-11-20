-- 00003_seed_annual_emissions.up.sql
-- Seed annual emissions for Stockholm (2025)

-- Ensure the city exists (upsert Stockholm) and capture its id
WITH city AS (
    INSERT INTO Cities (city_name, total_budget_co2e, budget_set_year)
    VALUES ('Stockholm', 25000000, 2025)
    ON CONFLICT (city_name) DO UPDATE SET total_budget_co2e = EXCLUDED.total_budget_co2e, budget_set_year = EXCLUDED.budget_set_year
    RETURNING city_id
)
INSERT INTO AnnualEmissions (city_id, activity_id, year, annual_emissions_co2e, data_source)
VALUES
    ((SELECT city_id FROM city), (SELECT activity_id FROM ActivityCategories WHERE category_name = 'Transport' LIMIT 1), 2025, 1200000, 'seed'),
    ((SELECT city_id FROM city), (SELECT activity_id FROM ActivityCategories WHERE category_name = 'Buildings' LIMIT 1), 2025, 1500000, 'seed'),
    ((SELECT city_id FROM city), (SELECT activity_id FROM ActivityCategories WHERE category_name = 'Industry' LIMIT 1), 2025, 800000, 'seed'),
    ((SELECT city_id FROM city), (SELECT activity_id FROM ActivityCategories WHERE category_name = 'Waste' LIMIT 1), 2025, 300000, 'seed')
ON CONFLICT (city_id, activity_id, year) DO UPDATE
    SET annual_emissions_co2e = EXCLUDED.annual_emissions_co2e, data_source = EXCLUDED.data_source;
