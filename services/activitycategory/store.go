package activitycategory

import (
	"context"
	"database/sql"

	"github.com/AboLojy/Carbon-Budget-Visualiser/types"
)

// Store defines persistence operations for activity categories.
type ActivityStore interface {
	GetAll(ctx context.Context) ([]types.Activity, error)
}

// Store implements Store using *sql.DB.
type Store struct {
	db *sql.DB
}

// NewStore creates a Store backed by the provided DB.
func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// GetAll returns all activity categories from ActivityCategories table.
func (s *Store) GetAll(ctx context.Context) ([]types.Activity, error) {
	const query = `SELECT activity_id, category_name, description FROM ActivityCategories ORDER BY category_name`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []types.Activity
	for rows.Next() {
		var a types.Activity
		if err := rows.Scan(&a.ActivityID, &a.CategoryName, &a.Description); err != nil {
			return nil, err
		}
		out = append(out, a)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return out, nil
}
