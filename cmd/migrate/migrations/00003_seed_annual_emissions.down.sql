-- 00003_seed_annual_emissions.down.sql
-- Rollback seed for annual emissions (Stockholm, 2025)

DELETE FROM AnnualEmissions
WHERE year = 2025
  AND city_id = (
    SELECT city_id FROM Cities WHERE city_name = 'Stockholm' LIMIT 1
  );

-- (Optional) remove the city if you want; commented out by default
-- DELETE FROM Cities WHERE city_name = 'Stockholm';
