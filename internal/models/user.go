package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"type:varchar(100);not null" validate:"required,min=2,max=100"`
	Email     string         `json:"email" gorm:"type:varchar(100);uniqueIndex;not null" validate:"required,email"`
	Phone     string         `json:"phone" gorm:"type:varchar(20);not null" validate:"required,min=10,max=15"`
	Address   string         `json:"address" gorm:"type:text"`
	Role      string         `json:"role" gorm:"type:varchar(20);default:member" validate:"oneof=admin librarian member"`
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	Borrows []Borrow `json:"borrows,omitempty" gorm:"foreignKey:UserID"`
}

type CreateUserRequest struct {
	Name    string `json:"name" validate:"required,min=2,max=100"`
	Email   string `json:"email" validate:"required,email"`
	Phone   string `json:"phone" validate:"required,min=10,max=15"`
	Address string `json:"address"`
	Role    string `json:"role" validate:"oneof=admin librarian member"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" validate:"min=2,max=100"`
	Email    string `json:"email" validate:"email"`
	Phone    string `json:"phone" validate:"min=10,max=15"`
	Address  string `json:"address"`
	Role     string `json:"role" validate:"oneof=admin librarian member"`
	IsActive *bool  `json:"is_active"`
}
