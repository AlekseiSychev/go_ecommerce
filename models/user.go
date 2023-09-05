package models

import "time"

type Users struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"size:255 not null"`
	Email     string    `json:"email" gorm:"size:255 unique not null"`
	Phone     string    `json:"phone" gorm:"size:20 unique not null"`
	Password  string    `json:"password" gorm:"size:255 not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
