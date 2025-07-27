package models

import (
	"time"

	"gorm.io/gorm"
)

type BorrowStatus string

const (
	StatusBorrowed BorrowStatus = "borrowed"
	StatusReturned BorrowStatus = "returned"
	StatusOverdue  BorrowStatus = "overdue"
	StatusLost     BorrowStatus = "lost"
)

type Borrow struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	UserID     uint           `json:"user_id" gorm:"not null" validate:"required"`
	BookID     uint           `json:"book_id" gorm:"not null" validate:"required"`
	BorrowDate time.Time      `json:"borrow_date" gorm:"not null default:current_timestamp" validate:"required"` // not null and default current timestamp
	DueDate    time.Time      `json:"due_date" gorm:"not null" validate:"required"`
	ReturnDate *time.Time     `json:"return_date"`
	Status     BorrowStatus   `json:"status" gorm:"default:borrowed" validate:"oneof=borrowed returned overdue lost"`
	Fine       float64        `json:"fine" gorm:"default:0"`
	Notes      string         `json:"notes" gorm:"type:text"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	User User `json:"user" gorm:"foreignKey:UserID"`
	Book Book `json:"book" gorm:"foreignKey:BookID"`
}

type CreateBorrowRequest struct {
	UserID  uint      `json:"user_id" validate:"required"`
	BookID  uint      `json:"book_id" validate:"required"`
	DueDate time.Time `json:"due_date" validate:"required"`
	Notes   string    `json:"notes"`
}

type UpdateBorrowRequest struct {
	DueDate    *time.Time    `json:"due_date"`
	ReturnDate *time.Time    `json:"return_date"`
	Status     *BorrowStatus `json:"status" validate:"omitempty,oneof=borrowed returned overdue lost"`
	Fine       *float64      `json:"fine" validate:"omitempty,min=0"`
	Notes      string        `json:"notes"`
}

type ReturnBookRequest struct {
	Fine  float64 `json:"fine" validate:"min=0"`
	Notes string  `json:"notes"`
}
