package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/yooerizkilab/library-system/internal/models"
	"github.com/yooerizkilab/library-system/internal/services"
	"github.com/yooerizkilab/library-system/pkg/response"
)

type BorrowHandler struct {
	borrowService services.BorrowService
}

func NewBorrowHandler(borrowService services.BorrowService) *BorrowHandler {
	return &BorrowHandler{
		borrowService: borrowService,
	}
}

func (h *BorrowHandler) BorrowBook(c *fiber.Ctx) error {
	var req models.CreateBorrowRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	borrow, err := h.borrowService.BorrowBook(&req)
	if err != nil {
		return response.BadRequest(c, "Failed to borrow book", err.Error())
	}

	return response.Created(c, "Book borrowed successfully", borrow)
}

func (h *BorrowHandler) ReturnBook(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return response.BadRequest(c, "Invalid borrow ID", err.Error())
	}

	var req models.ReturnBookRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	borrow, err := h.borrowService.ReturnBook(uint(id), &req)
	if err != nil {
		if err.Error() == "borrow record not found" {
			return response.NotFound(c, "Borrow record not found")
		}
		return response.BadRequest(c, "Failed to return book", err.Error())
	}

	return response.Success(c, "Book returned successfully", borrow)
}

func (h *BorrowHandler) GetAllBorrows(c *fiber.Ctx) error {
	borrows, err := h.borrowService.GetAllBorrows()
	if err != nil {
		return response.InternalServerError(c, "Failed to get borrows", err.Error())
	}

	return response.Success(c, "Borrows retrieved successfully", borrows)
}

func (h *BorrowHandler) GetBorrowByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return response.BadRequest(c, "Invalid borrow ID", err.Error())
	}

	borrow, err := h.borrowService.GetBorrowByID(uint(id))
	if err != nil {
		if err.Error() == "borrow record not found" {
			return response.NotFound(c, "Borrow record not found")
		}
		return response.InternalServerError(c, "Failed to get borrow", err.Error())
	}

	return response.Success(c, "Borrow retrieved successfully", borrow)
}

func (h *BorrowHandler) GetBorrowsByUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("userId"), 10, 32)
	if err != nil {
		return response.BadRequest(c, "Invalid user ID", err.Error())
	}

	borrows, err := h.borrowService.GetBorrowsByUser(uint(id))
	if err != nil {
		if err.Error() == "user not found" {
			return response.NotFound(c, "User not found")
		}
		return response.InternalServerError(c, "Failed to get user borrows", err.Error())
	}

	return response.Success(c, "User borrows retrieved successfully", borrows)
}

func (h *BorrowHandler) GetBorrowsByBook(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("bookId"), 10, 32)
	if err != nil {
		return response.BadRequest(c, "Invalid book ID", err.Error())
	}

	borrows, err := h.borrowService.GetBorrowsByBook(uint(id))
	if err != nil {
		if err.Error() == "book not found" {
			return response.NotFound(c, "Book not found")
		}
		return response.InternalServerError(c, "Failed to get book borrows", err.Error())
	}

	return response.Success(c, "Book borrows retrieved successfully", borrows)
}

func (h *BorrowHandler) UpdateBorrow(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return response.BadRequest(c, "Invalid borrow ID", err.Error())
	}

	var req models.UpdateBorrowRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	borrow, err := h.borrowService.UpdateBorrow(uint(id), &req)
	if err != nil {
		if err.Error() == "borrow record not found" {
			return response.NotFound(c, "Borrow record not found")
		}
		return response.BadRequest(c, "Failed to update borrow", err.Error())
	}

	return response.Success(c, "Borrow updated successfully", borrow)
}

func (h *BorrowHandler) GetActiveBorrows(c *fiber.Ctx) error {
	borrows, err := h.borrowService.GetActiveBorrows()
	if err != nil {
		return response.InternalServerError(c, "Failed to get active borrows", err.Error())
	}

	return response.Success(c, "Active borrows retrieved successfully", borrows)
}

func (h *BorrowHandler) GetOverdueBorrows(c *fiber.Ctx) error {
	borrows, err := h.borrowService.GetOverdueBorrows()
	if err != nil {
		return response.InternalServerError(c, "Failed to get overdue borrows", err.Error())
	}

	return response.Success(c, "Overdue borrows retrieved successfully", borrows)
}

func (h *BorrowHandler) GetBorrowHistory(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("userId"), 10, 32)
	if err != nil {
		return response.BadRequest(c, "Invalid user ID", err.Error())
	}

	borrows, err := h.borrowService.GetBorrowHistory(uint(id))
	if err != nil {
		if err.Error() == "user not found" {
			return response.NotFound(c, "User not found")
		}
		return response.InternalServerError(c, "Failed to get borrow history", err.Error())
	}

	return response.Success(c, "Borrow history retrieved successfully", borrows)
}
