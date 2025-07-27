package services

import (
	"errors"

	"github.com/yooerizkilab/library-system/internal/models"
	"github.com/yooerizkilab/library-system/internal/repositories"
	"gorm.io/gorm"
)

type BookService interface {
	CreateBook(req *models.CreateBookRequest) (*models.Book, error)
	GetAllBooks() ([]models.Book, error)
	GetBookByID(id uint) (*models.Book, error)
	UpdateBook(id uint, req *models.UpdateBookRequest) (*models.Book, error)
	DeleteBook(id uint) error
	SearchBooks(query string) ([]models.Book, error)
	GetBooksByCategory(category string) ([]models.Book, error)
	GetAvailableBooks() ([]models.Book, error)
}

type bookService struct {
	bookRepo repositories.BookRepository
}

func NewBookService(bookRepo repositories.BookRepository) BookService {
	return &bookService{
		bookRepo: bookRepo,
	}
}

func (s *bookService) CreateBook(req *models.CreateBookRequest) (*models.Book, error) {
	// Check if ISBN already exists
	existingBook, err := s.bookRepo.GetByISBN(req.ISBN)
	if err == nil && existingBook != nil {
		return nil, errors.New("book with this ISBN already exists")
	}

	// Set default values
	if req.Language == "" {
		req.Language = "Indonesian"
	}
	if req.Stock == 0 {
		req.Stock = 1
	}

	book := &models.Book{
		Title:       req.Title,
		Author:      req.Author,
		ISBN:        req.ISBN,
		Publisher:   req.Publisher,
		Category:    req.Category,
		Language:    req.Language,
		Pages:       req.Pages,
		PublishYear: req.PublishYear,
		Stock:       req.Stock,
		Available:   req.Stock, // Initially all books are available
		Description: req.Description,
		Location:    req.Location,
		IsActive:    true,
	}

	err = s.bookRepo.Create(book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *bookService) GetAllBooks() ([]models.Book, error) {
	return s.bookRepo.GetAll()
}

func (s *bookService) GetBookByID(id uint) (*models.Book, error) {
	book, err := s.bookRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}
	return book, nil
}

func (s *bookService) UpdateBook(id uint, req *models.UpdateBookRequest) (*models.Book, error) {
	book, err := s.bookRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}

	// Check if ISBN is being changed and if it already exists
	if req.ISBN != "" && req.ISBN != book.ISBN {
		existingBook, err := s.bookRepo.GetByISBN(req.ISBN)
		if err == nil && existingBook != nil {
			return nil, errors.New("book with this ISBN already exists")
		}
	}

	// Store original stock to calculate available books
	originalStock := book.Stock
	borrowedBooks := originalStock - book.Available

	// Update fields
	if req.Title != "" {
		book.Title = req.Title
	}
	if req.Author != "" {
		book.Author = req.Author
	}
	if req.ISBN != "" {
		book.ISBN = req.ISBN
	}
	if req.Publisher != "" {
		book.Publisher = req.Publisher
	}
	if req.Category != "" {
		book.Category = req.Category
	}
	if req.Language != "" {
		book.Language = req.Language
	}
	if req.Pages > 0 {
		book.Pages = req.Pages
	}
	if req.PublishYear > 0 {
		book.PublishYear = req.PublishYear
	}
	if req.Stock >= 0 {
		book.Stock = req.Stock
		// Recalculate available books
		book.Available = req.Stock - borrowedBooks
		if book.Available < 0 {
			book.Available = 0
		}
	}
	if req.Description != "" {
		book.Description = req.Description
	}
	if req.Location != "" {
		book.Location = req.Location
	}
	if req.IsActive != nil {
		book.IsActive = *req.IsActive
	}

	err = s.bookRepo.Update(book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *bookService) DeleteBook(id uint) error {
	book, err := s.bookRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("book not found")
		}
		return err
	}

	// Check if book has active borrows
	if len(book.Borrows) > 0 {
		for _, borrow := range book.Borrows {
			if borrow.Status == models.StatusBorrowed {
				return errors.New("cannot delete book with active borrows")
			}
		}
	}

	return s.bookRepo.Delete(id)
}

func (s *bookService) SearchBooks(query string) ([]models.Book, error) {
	return s.bookRepo.Search(query)
}

func (s *bookService) GetBooksByCategory(category string) ([]models.Book, error) {
	return s.bookRepo.GetByCategory(category)
}

func (s *bookService) GetAvailableBooks() ([]models.Book, error) {
	return s.bookRepo.GetAvailableBooks()
}
