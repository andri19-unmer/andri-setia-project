package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type ProductRepository interface {
	Fetch(ctx context.Context) ([]Product, error)
	GetByID(ctx context.Context, id uint) (Product, error)
	Store(ctx context.Context, product *Product) error
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id uint) error
}

type ProductUsecase interface {
	Fetch(ctx context.Context) ([]Product, error)
	GetByID(ctx context.Context, id uint) (Product, error)
	Store(ctx context.Context, product *Product) error
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id uint) error
}
