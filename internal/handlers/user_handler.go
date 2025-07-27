package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/yooerizkilab/library-system/internal/models"
	"github.com/yooerizkilab/library-system/internal/services"
	"github.com/yooerizkilab/library-system/pkg/response"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	user, err := h.userService.CreateUser(&req)
	if err != nil {
		return response.BadRequest(c, "Failed to create user", err.Error())
	}

	return response.Created(c, "User created successfully", user)
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		return response.InternalServerError(c, "Failed to get users", err.Error())
	}

	return response.Success(c, "Users retrieved successfully", users)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return response.BadRequest(c, "Invalid user ID", err.Error())
	}

	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		if err.Error() == "user not found" {
			return response.NotFound(c, "User not found")
		}
		return response.InternalServerError(c, "Failed to get user", err.Error())
	}

	return response.Success(c, "User retrieved successfully", user)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return response.BadRequest(c, "Invalid user ID", err.Error())
	}

	var req models.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	user, err := h.userService.UpdateUser(uint(id), &req)
	if err != nil {
		if err.Error() == "user not found" {
			return response.NotFound(c, "User not found")
		}
		return response.BadRequest(c, "Failed to update user", err.Error())
	}

	return response.Success(c, "User updated successfully", user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return response.BadRequest(c, "Invalid user ID", err.Error())
	}

	err = h.userService.DeleteUser(uint(id))
	if err != nil {
		if err.Error() == "user not found" {
			return response.NotFound(c, "User not found")
		}
		return response.BadRequest(c, "Failed to delete user", err.Error())
	}

	return response.Success(c, "User deleted successfully", nil)
}

func (h *UserHandler) SearchUsers(c *fiber.Ctx) error {
	query := c.Query("q")
	if query == "" {
		return response.BadRequest(c, "Search query is required", nil)
	}

	users, err := h.userService.SearchUsers(query)
	if err != nil {
		return response.InternalServerError(c, "Failed to search users", err.Error())
	}

	return response.Success(c, "Users search completed", users)
}
