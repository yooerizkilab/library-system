package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/yooerizkilab/library-system/internal/models"
	"github.com/yooerizkilab/library-system/internal/services"
	"github.com/yooerizkilab/library-system/pkg/response"
)

type BookHandler struct {
	bookService services.BookService
}

func NewBookHandler(bookService services.BookService) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}

func (h *BookHandler) CreateBook(c *fiber.Ctx) error {
	var req models.CreateBookRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	book, err := h.bookService.CreateBook(&req)
	if err != nil {
		return response.BadRequest(c, "Failed to create book", err.Error())
	}

	return response.Created(c, "Book created successfully", book)
}

func (h *BookHandler) GetAllBooks(c *fiber.Ctx) error {
	books, err := h.bookService.GetAllBooks()
	if err != nil {
		return response.InternalServerError(c, "Failed to get books", err.Error())
	}

	return response.Success(c, "Books retrieved successfully", books)
}

func (h *BookHandler) GetBookByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return response.BadRequest(c, "Invalid book ID", err.Error())
	}

	book, err := h.bookService.GetBookByID(uint(id))
	if err != nil {
		if err.Error() == "book not found" {
			return response.NotFound(c, "Book not found")
		}
		return response.InternalServerError(c, "Failed to get book", err.Error())
	}

	return response.Success(c, "Book retrieved successfully", book)
}

func (h *BookHandler) UpdateBook(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return response.BadRequest(c, "Invalid book ID", err.Error())
	}

	var req models.UpdateBookRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	book, err := h.bookService.UpdateBook(uint(id), &req)
	if err != nil {
		if err.Error() == "book not found" {
			return response.NotFound(c, "Book not found")
		}
		return response.BadRequest(c, "Failed to update book", err.Error())
	}

	return response.Success(c, "Book updated successfully", book)
}

func (h *BookHandler) DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return response.BadRequest(c, "Invalid book ID", err.Error())
	}

	err = h.bookService.DeleteBook(uint(id))
	if err != nil {
		if err.Error() == "book not found" {
			return response.NotFound(c, "Book not found")
		}
		return response.BadRequest(c, "Failed to delete book", err.Error())
	}

	return response.Success(c, "Book deleted successfully", nil)
}

func (h *BookHandler) SearchBooks(c *fiber.Ctx) error {
	query := c.Query("q")
	if query == "" {
		return response.BadRequest(c, "Search query is required", nil)
	}

	books, err := h.bookService.SearchBooks(query)
	if err != nil {
		return response.InternalServerError(c, "Failed to search books", err.Error())
	}

	return response.Success(c, "Books search completed", books)
}

func (h *BookHandler) GetBooksByCategory(c *fiber.Ctx) error {
	category := c.Params("category")
	if category == "" {
		return response.BadRequest(c, "Category is required", nil)
	}

	books, err := h.bookService.GetBooksByCategory(category)
	if err != nil {
		return response.InternalServerError(c, "Failed to get books by category", err.Error())
	}

	return response.Success(c, "Books retrieved successfully", books)
}

func (h *BookHandler) GetAvailableBooks(c *fiber.Ctx) error {
	books, err := h.bookService.GetAvailableBooks()
	if err != nil {
		return response.InternalServerError(c, "Failed to get available books", err.Error())
	}

	return response.Success(c, "Available books retrieved successfully", books)
}
