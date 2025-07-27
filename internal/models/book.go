package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"type:varchar(200);not null" validate:"required,min=1,max=200"`
	Author      string         `json:"author" gorm:"type:varchar(100);not null" validate:"required,min=1,max=100"`
	ISBN        string         `json:"isbn" gorm:"type:varchar(20);uniqueIndex" validate:"required,min=10,max=17"`
	Publisher   string         `json:"publisher" gorm:"type:varchar(100)" validate:"max=100"`
	Category    string         `json:"category" gorm:"type:varchar(50);not null" validate:"required,max=50"`
	Language    string         `json:"language" gorm:"type:varchar(30);default:Indonesian" validate:"max=30"`
	Pages       int            `json:"pages" validate:"min=1"`
	PublishYear int            `json:"publish_year" validate:"min=1000,max=2100"`
	Stock       int            `json:"stock" gorm:"default:1" validate:"min=0"`
	Available   int            `json:"available" gorm:"default:1" validate:"min=0"`
	Description string         `json:"description" gorm:"type:text"`
	Location    string         `json:"location" gorm:"type:varchar(50)" validate:"max=50"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	Borrows []Borrow `json:"borrows,omitempty" gorm:"foreignKey:BookID"`
}

type CreateBookRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=200"`
	Author      string `json:"author" validate:"required,min=1,max=100"`
	ISBN        string `json:"isbn" validate:"required,min=10,max=17"`
	Publisher   string `json:"publisher" validate:"max=100"`
	Category    string `json:"category" validate:"required,max=50"`
	Language    string `json:"language" validate:"max=30"`
	Pages       int    `json:"pages" validate:"min=1"`
	PublishYear int    `json:"publish_year" validate:"min=1000,max=2100"`
	Stock       int    `json:"stock" validate:"min=1"`
	Description string `json:"description"`
	Location    string `json:"location" validate:"max=50"`
}

type UpdateBookRequest struct {
	Title       string `json:"title" validate:"min=1,max=200"`
	Author      string `json:"author" validate:"min=1,max=100"`
	ISBN        string `json:"isbn" validate:"min=10,max=17"`
	Publisher   string `json:"publisher" validate:"max=100"`
	Category    string `json:"category" validate:"max=50"`
	Language    string `json:"language" validate:"max=30"`
	Pages       int    `json:"pages" validate:"min=1"`
	PublishYear int    `json:"publish_year" validate:"min=1000,max=2100"`
	Stock       int    `json:"stock" validate:"min=0"`
	Description string `json:"description"`
	Location    string `json:"location" validate:"max=50"`
	IsActive    *bool  `json:"is_active"`
}
