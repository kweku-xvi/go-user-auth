package models

import "time"

type User struct {
	ID        string    `gorm:"column:id;primaryKey" json:"id"`
	FirstName string    `gorm:"column:first_name" json:"first_name"`
	LastName  string    `gorm:"column:last_name" json:"last_name"`
	Email     string    `gorm:"column:email;notnull;unique" json:"email"`
	Password  string    `gorm:"column:password;notnull" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
