package repositories

import (
	"github.com/yooerizkilab/library-system/internal/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book *models.Book) error
	GetAll() ([]models.Book, error)
	GetByID(id uint) (*models.Book, error)
	GetByISBN(isbn string) (*models.Book, error)
	Update(book *models.Book) error
	Delete(id uint) error
	Search(query string) ([]models.Book, error)
	GetByCategory(category string) ([]models.Book, error)
	GetAvailableBooks() ([]models.Book, error)
	UpdateStock(id uint, stock, available int) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) Create(book *models.Book) error {
	// Set available same as stock initially
	book.Available = book.Stock
	return r.db.Create(book).Error
}

func (r *bookRepository) GetAll() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Where("is_active = ?", true).Find(&books).Error
	return books, err
}

func (r *bookRepository) GetByID(id uint) (*models.Book, error) {
	var book models.Book
	err := r.db.Preload("Borrows").First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *bookRepository) GetByISBN(isbn string) (*models.Book, error) {
	var book models.Book
	err := r.db.Where("isbn = ?", isbn).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *bookRepository) Update(book *models.Book) error {
	return r.db.Save(book).Error
}

func (r *bookRepository) Delete(id uint) error {
	return r.db.Delete(&models.Book{}, id).Error
}

func (r *bookRepository) Search(query string) ([]models.Book, error) {
	var books []models.Book
	err := r.db.Where("is_active = ? AND (title LIKE ? OR author LIKE ? OR isbn LIKE ? OR category LIKE ?)",
		true, "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%").Find(&books).Error
	return books, err
}

func (r *bookRepository) GetByCategory(category string) ([]models.Book, error) {
	var books []models.Book
	err := r.db.Where("category = ? AND is_active = ?", category, true).Find(&books).Error
	return books, err
}

func (r *bookRepository) GetAvailableBooks() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Where("available > 0 AND is_active = ?", true).Find(&books).Error
	return books, err
}

func (r *bookRepository) UpdateStock(id uint, stock, available int) error {
	return r.db.Model(&models.Book{}).Where("id = ?", id).Updates(map[string]interface{}{
		"stock":     stock,
		"available": available,
	}).Error
}
