# Project Carbon Budget Visualiser 

Help cities visualize their carbon budgets and plan for exhaustion

## Architecture

### Layered Design

The project follows a **service-oriented layered architecture**:

- **`cmd/`**: Application entry points (API server, database migrations).
- **`services/`**: Service packages, each with its own store (data access) and routes (HTTP handlers).
  - `emission/`: City creation and management.
  - `activitycategory/`: Activity category lookups (seeded data).
  - `annualemission/`: Annual emission records (create, list, delete).
  - `budgetcalculator/`: Budget consumption projections (calculates remaining budget year-by-year).
- **`types/`**: Domain models (DTOs) shared across services.
- **`db/`**: Database connection and migrations.
- **`config/`**: Environment and configuration loading.

### Technical Choices

- **`net/http` (stdlib)**: Built-in HTTP server and routing. No external frameworks for simplicity.
- **`go-playground/validator`**: Struct-level validation with tags (e.g., `validate:"required,gt=0"`).
- **`golang-migrate`**: SQL-based schema versioning and seeding via `.up.sql` and `.down.sql` files.
- **Store Interface Pattern**: Each service defines a `Store` interface for data access, enabling easy testing and swapping implementations.

### Data Model

- **Cities**: city name, total CO₂ budget, budget year.
- **ActivityCategories**: predefined emission sources (Transport, Buildings, Industry, Waste, etc.).
- **AnnualEmissions**: emissions per activity per city per year.
- **Budget Projection**: calculated remaining budget after subtracting annual emissions year-by-year.



## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

## MakeFile

Create DB container 

```bash
make docker-run
```
Create database and seed it  

```bash
make db-up
```

Run the application
```bash
make run
```

Remove database 

```bash
make db-down
```

Shutdown DB Container

```bash
make docker-down
```

## API Routes

Short descriptions and example curl commands for available API endpoints (assumes server at `http://localhost:8080`):

- **Create City**: create a new city record.

```bash
curl -X POST "http://localhost:8080/api/v1/city" \
	-H "Content-Type: application/json" \
	-d '{"city_name":"Stockholm","total_budget_co2e":2500000,"budget_set_year":2025}'
```

- **Create City Budget** (budget calculator): create/update a city's budget.

```bash
curl -X POST "http://localhost:8080/api/v1/citybudget" \
	-H "Content-Type: application/json" \
	-d '{"city_name":"Cairo","total_budget_co2e":55000000,"budget_set_year":2022}'
```

- **Get City Budget Projection**: compute year-by-year remaining budget for a city starting a given year.

```bash
curl -X GET "http://localhost:8080/api/v1/citybudget/2/2022"
```

- **List Activity Categories**: returns seeded activity categories (Transport, Buildings, etc.).

```bash
curl -X GET "http://localhost:8080/api/v1/activities"
```

- **Create Annual Emission**: insert an annual emission row for a city/activity/year.

```bash
curl -X POST "http://localhost:8080/api/v1/annualemissions" \
	-H "Content-Type: application/json" \
	-d '{"city_id":1,"activity_id":3,"year":2025,"annual_emissions_co2e":1200000,"data_source":"seed"}'
```

- **List Annual Emissions**: list all annual emission rows.

```bash
curl -X GET "http://localhost:8080/api/v1/annualemissions"
```

- **Delete Annual Emission**: delete by emission id (query param) or path if supported.

```bash
curl -X DELETE "http://localhost:8080/api/v1/annualemissions?id=123"
```

