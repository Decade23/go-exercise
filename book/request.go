package book

import (
	jsoniter "github.com/json-iterator/go"
	"time"
)

type BookRequest struct {
	ID        int             `json:"id" gorm:"primaryKey"`
	Title     string          `json:"title" binding:"required"`
	Discount  int             `json:"discount"`
	Price     jsoniter.Number `json:"price" binding:"required,number"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
