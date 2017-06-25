package models

import (
	"time"
)

// Todo todo item information.
type Todo struct {
	ID       string     `json:"id" bson:"_id"`
	UserID   string     `json:"userId" bson:"userId"`
	Content  string     `json:"content" bson:"content"`
	Category string     `json:"category" bson:"category"`
	Deadline *time.Time `json:"deadline" bson:"deadline"`
	Priority int        `json:"priority" bson:"priority"`
	Done     bool       `json:"done" bson:"done"`
}
