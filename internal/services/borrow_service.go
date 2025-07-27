package services

import (
	"errors"
	"time"

	"github.com/yooerizkilab/library-system/internal/models"
	"github.com/yooerizkilab/library-system/internal/repositories"
	"gorm.io/gorm"
)

type BorrowService interface {
	BorrowBook(req *models.CreateBorrowRequest) (*models.Borrow, error)
	ReturnBook(id uint, req *models.ReturnBookRequest) (*models.Borrow, error)
	GetAllBorrows() ([]models.Borrow, error)
	GetBorrowByID(id uint) (*models.Borrow, error)
	GetBorrowsByUser(userID uint) ([]models.Borrow, error)
	GetBorrowsByBook(bookID uint) ([]models.Borrow, error)
	UpdateBorrow(id uint, req *models.UpdateBorrowRequest) (*models.Borrow, error)
	GetActiveBorrows() ([]models.Borrow, error)
	GetOverdueBorrows() ([]models.Borrow, error)
	GetBorrowHistory(userID uint) ([]models.Borrow, error)
	UpdateOverdueStatus() error
}

type borrowService struct {
	borrowRepo repositories.BorrowRepository
	userRepo   repositories.UserRepository
	bookRepo   repositories.BookRepository
}

func NewBorrowService(
	borrowRepo repositories.BorrowRepository,
	userRepo repositories.UserRepository,
	bookRepo repositories.BookRepository,
) BorrowService {
	return &borrowService{
		borrowRepo: borrowRepo,
		userRepo:   userRepo,
		bookRepo:   bookRepo,
	}
}

func (s *borrowService) BorrowBook(req *models.CreateBorrowRequest) (*models.Borrow, error) {
	// Check if user exists and is active
	user, err := s.userRepo.GetByID(req.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	if !user.IsActive {
		return nil, errors.New("user is not active")
	}

	// Check if book exists and is available
	book, err := s.bookRepo.GetByID(req.BookID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}
	if !book.IsActive {
		return nil, errors.New("book is not active")
	}
	if book.Available <= 0 {
		return nil, errors.New("book is not available for borrowing")
	}

	// Check if user already has an active borrow for this book
	activeBorrow, err := s.borrowRepo.CheckActiveUserBorrow(req.UserID, req.BookID)
	if err == nil && activeBorrow != nil {
		return nil, errors.New("user already has an active borrow for this book")
	}

	// Create borrow record
	borrow := &models.Borrow{
		UserID:     req.UserID,
		BookID:     req.BookID,
		BorrowDate: time.Now(),
		DueDate:    req.DueDate,
		Status:     models.StatusBorrowed,
		Notes:      req.Notes,
	}

	err = s.borrowRepo.Create(borrow)
	if err != nil {
		return nil, err
	}

	// Update book availability
	err = s.bookRepo.UpdateStock(req.BookID, book.Stock, book.Available-1)
	if err != nil {
		return nil, err
	}

	// Get the created borrow with relations
	createdBorrow, err := s.borrowRepo.GetByID(borrow.ID)
	if err != nil {
		return nil, err
	}

	return createdBorrow, nil
}

func (s *borrowService) ReturnBook(id uint, req *models.ReturnBookRequest) (*models.Borrow, error) {
	borrow, err := s.borrowRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("borrow record not found")
		}
		return nil, err
	}

	if borrow.Status != models.StatusBorrowed {
		return nil, errors.New("book is not currently borrowed")
	}

	// Update borrow record
	now := time.Now()
	borrow.ReturnDate = &now
	borrow.Status = models.StatusReturned
	borrow.Fine = req.Fine
	borrow.Notes = req.Notes

	err = s.borrowRepo.Update(borrow)
	if err != nil {
		return nil, err
	}

	// Update book availability
	book, err := s.bookRepo.GetByID(borrow.BookID)
	if err != nil {
		return nil, err
	}

	err = s.bookRepo.UpdateStock(borrow.BookID, book.Stock, book.Available+1)
	if err != nil {
		return nil, err
	}

	return borrow, nil
}

func (s *borrowService) GetAllBorrows() ([]models.Borrow, error) {
	return s.borrowRepo.GetAll()
}

func (s *borrowService) GetBorrowByID(id uint) (*models.Borrow, error) {
	borrow, err := s.borrowRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("borrow record not found")
		}
		return nil, err
	}
	return borrow, nil
}

func (s *borrowService) GetBorrowsByUser(userID uint) ([]models.Borrow, error) {
	// Check if user exists
	_, err := s.userRepo.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return s.borrowRepo.GetByUserID(userID)
}

func (s *borrowService) GetBorrowsByBook(bookID uint) ([]models.Borrow, error) {
	// Check if book exists
	_, err := s.bookRepo.GetByID(bookID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}

	return s.borrowRepo.GetByBookID(bookID)
}

func (s *borrowService) UpdateBorrow(id uint, req *models.UpdateBorrowRequest) (*models.Borrow, error) {
	borrow, err := s.borrowRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("borrow record not found")
		}
		return nil, err
	}

	// Update fields
	if req.DueDate != nil {
		borrow.DueDate = *req.DueDate
	}
	if req.ReturnDate != nil {
		borrow.ReturnDate = req.ReturnDate
	}
	if req.Status != nil {
		borrow.Status = *req.Status
	}
	if req.Fine != nil {
		borrow.Fine = *req.Fine
	}
	if req.Notes != "" {
		borrow.Notes = req.Notes
	}

	err = s.borrowRepo.Update(borrow)
	if err != nil {
		return nil, err
	}

	return borrow, nil
}

func (s *borrowService) GetActiveBorrows() ([]models.Borrow, error) {
	return s.borrowRepo.GetActiveBorrows()
}

func (s *borrowService) GetOverdueBorrows() ([]models.Borrow, error) {
	return s.borrowRepo.GetOverdueBorrows()
}

func (s *borrowService) GetBorrowHistory(userID uint) ([]models.Borrow, error) {
	// Check if user exists
	_, err := s.userRepo.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return s.borrowRepo.GetBorrowHistory(userID)
}

func (s *borrowService) UpdateOverdueStatus() error {
	overdueBorrows, err := s.borrowRepo.GetOverdueBorrows()
	if err != nil {
		return err
	}

	for _, borrow := range overdueBorrows {
		if borrow.Status == models.StatusBorrowed {
			borrow.Status = models.StatusOverdue
			err = s.borrowRepo.Update(&borrow)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
