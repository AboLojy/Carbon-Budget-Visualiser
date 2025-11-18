-- SQL in section 'Down' is executed when this migration is rolled back
-- 00001_create_initial_schema.down.sql

DROP TABLE IF EXISTS AnnualEmissions;
DROP TABLE IF EXISTS ActivityCategories;
DROP TABLE IF EXISTS Cities;