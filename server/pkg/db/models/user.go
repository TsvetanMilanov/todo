package models

// User the db model for the users collection
type User struct {
	ID       string `bson:"_id"`
	Username string
	Password string
	Todos    []string
	Roles    []string
}
