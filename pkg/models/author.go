package models

// struct for database schema
type AuthorDetail struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	FirstName string
	LastName  string
}
