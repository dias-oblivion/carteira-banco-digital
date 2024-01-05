package models

type User struct {
	ID         int64
	Name       string
	Document   string
	Email      string
	Password   string
	Role       string
	Balance    float64
	Created_at string
}
