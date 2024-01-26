package models

// struct for database schema
type UserDetail struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Username string
	Password string
}
