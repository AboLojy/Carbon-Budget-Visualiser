-- 00002_seed_activity_categories.up.sql

INSERT INTO ActivityCategories (category_name, description) VALUES
('Transport', 'Emissions from road, rail, air, and marine travel.'),
('Buildings', 'Emissions from heating, cooling, and electricity use in residential and commercial properties.'),
('Waste', 'Emissions from landfills, waste treatment, and recycling processes.'),
('Industry', 'Emissions from manufacturing and industrial processes.'),
('Electricity Generation', 'Emissions from local power plants supplying the city.'),
('Agriculture/Land Use', 'Emissions related to urban farming and changes in land use.')
ON CONFLICT (category_name) DO NOTHING;