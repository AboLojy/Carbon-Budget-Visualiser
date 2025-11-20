package types

// Activity represents an activity category in the DB
type Activity struct {
	ActivityID   int    `json:"activity_id" db:"activity_id"`
	CategoryName string `json:"category_name" db:"category_name"`
	Description  string `json:"description,omitempty" db:"description"`
}
