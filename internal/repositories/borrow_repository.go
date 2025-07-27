package repositories

import (
	"time"

	"github.com/yooerizkilab/library-system/internal/models"
	"gorm.io/gorm"
)

type BorrowRepository interface {
	Create(borrow *models.Borrow) error
	GetAll() ([]models.Borrow, error)
	GetByID(id uint) (*models.Borrow, error)
	GetByUserID(userID uint) ([]models.Borrow, error)
	GetByBookID(bookID uint) ([]models.Borrow, error)
	Update(borrow *models.Borrow) error
	Delete(id uint) error
	GetActiveBorrows() ([]models.Borrow, error)
	GetOverdueBorrows() ([]models.Borrow, error)
	GetBorrowHistory(userID uint) ([]models.Borrow, error)
	CheckActiveUserBorrow(userID, bookID uint) (*models.Borrow, error)
}

type borrowRepository struct {
	db *gorm.DB
}

func NewBorrowRepository(db *gorm.DB) BorrowRepository {
	return &borrowRepository{db: db}
}

func (r *borrowRepository) Create(borrow *models.Borrow) error {
	return r.db.Create(borrow).Error
}

func (r *borrowRepository) GetAll() ([]models.Borrow, error) {
	var borrows []models.Borrow
	err := r.db.Preload("User").Preload("Book").Find(&borrows).Error
	return borrows, err
}

func (r *borrowRepository) GetByID(id uint) (*models.Borrow, error) {
	var borrow models.Borrow
	err := r.db.Preload("User").Preload("Book").First(&borrow, id).Error
	if err != nil {
		return nil, err
	}
	return &borrow, nil
}

func (r *borrowRepository) GetByUserID(userID uint) ([]models.Borrow, error) {
	var borrows []models.Borrow
	err := r.db.Preload("User").Preload("Book").Where("user_id = ?", userID).Find(&borrows).Error
	return borrows, err
}

func (r *borrowRepository) GetByBookID(bookID uint) ([]models.Borrow, error) {
	var borrows []models.Borrow
	err := r.db.Preload("User").Preload("Book").Where("book_id = ?", bookID).Find(&borrows).Error
	return borrows, err
}

func (r *borrowRepository) Update(borrow *models.Borrow) error {
	return r.db.Save(borrow).Error
}

func (r *borrowRepository) Delete(id uint) error {
	return r.db.Delete(&models.Borrow{}, id).Error
}

func (r *borrowRepository) GetActiveBorrows() ([]models.Borrow, error) {
	var borrows []models.Borrow
	err := r.db.Preload("User").Preload("Book").Where("status = ?", models.StatusBorrowed).Find(&borrows).Error
	return borrows, err
}

func (r *borrowRepository) GetOverdueBorrows() ([]models.Borrow, error) {
	var borrows []models.Borrow
	now := time.Now()
	err := r.db.Preload("User").Preload("Book").Where("status = ? AND due_date < ?",
		models.StatusBorrowed, now).Find(&borrows).Error
	return borrows, err
}

func (r *borrowRepository) GetBorrowHistory(userID uint) ([]models.Borrow, error) {
	var borrows []models.Borrow
	err := r.db.Preload("User").Preload("Book").Where("user_id = ?", userID).
		Order("created_at DESC").Find(&borrows).Error
	return borrows, err
}

func (r *borrowRepository) CheckActiveUserBorrow(userID, bookID uint) (*models.Borrow, error) {
	var borrow models.Borrow
	err := r.db.Where("user_id = ? AND book_id = ? AND status = ?",
		userID, bookID, models.StatusBorrowed).First(&borrow).Error
	if err != nil {
		return nil, err
	}
	return &borrow, nil
}
