package routes

import (
	"github.com/yooerizkilab/library-system/internal/database"
	"github.com/yooerizkilab/library-system/internal/handlers"
	"github.com/yooerizkilab/library-system/internal/middleware"
	"github.com/yooerizkilab/library-system/internal/repositories"
	"github.com/yooerizkilab/library-system/internal/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Get database instance
	db := database.GetDB()

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	bookRepo := repositories.NewBookRepository(db)
	borrowRepo := repositories.NewBorrowRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepo)
	bookService := services.NewBookService(bookRepo)
	borrowService := services.NewBorrowService(borrowRepo, userRepo, bookRepo)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	bookHandler := handlers.NewBookHandler(bookService)
	borrowHandler := handlers.NewBorrowHandler(borrowService)

	// API version 1
	v1 := app.Group("/api/v1")

	// Public routes (no authentication required)
	auth := v1.Group("/auth")
	auth.Post("/login", userHandler.Login)
	auth.Post("/register", userHandler.CreateUser) // Public registration

	// Public book endpoints (read-only)
	publicBooks := v1.Group("/books")
	publicBooks.Get("/", bookHandler.GetAllBooks)
	publicBooks.Get("/search", bookHandler.SearchBooks)
	publicBooks.Get("/available", bookHandler.GetAvailableBooks)
	publicBooks.Get("/category/:category", bookHandler.GetBooksByCategory)
	publicBooks.Get("/:id", bookHandler.GetBookByID)

	// Protected routes (authentication required)
	protected := v1.Group("", middleware.AuthRequired())

	// User profile routes (authenticated users can access their own profile)
	profile := protected.Group("/profile")
	profile.Get("/", userHandler.GetProfile)
	profile.Put("/password", userHandler.ChangePassword)

	// User management routes (admin and librarian only)
	users := protected.Group("/users", middleware.RoleRequired("admin", "librarian"))
	users.Get("/", userHandler.GetAllUsers)
	users.Get("/search", userHandler.SearchUsers)
	users.Get("/:id", userHandler.GetUserByID)
	users.Put("/:id", userHandler.UpdateUser)
	users.Delete("/:id", middleware.RoleRequired("admin"), userHandler.DeleteUser) // Only admin can delete

	// Book management routes (admin and librarian only)
	bookManagement := protected.Group("/books/manage", middleware.RoleRequired("admin", "librarian"))
	bookManagement.Post("/", bookHandler.CreateBook)
	bookManagement.Put("/:id", bookHandler.UpdateBook)
	bookManagement.Delete("/:id", middleware.RoleRequired("admin"), bookHandler.DeleteBook) // Only admin can delete

	// Borrow routes
	borrows := protected.Group("/borrows")

	// All authenticated users can borrow books
	borrows.Post("/", borrowHandler.BorrowBook)

	// Users can view their own borrows, librarians and admins can view all
	borrows.Get("/", func(c *fiber.Ctx) error {
		userRole := c.Locals("user_role").(string)
		if userRole == "member" {
			// Members can only see their own borrows
			// userID := c.Locals("user_id").(uint)
			return borrowHandler.GetBorrowsByUser(c)
		}
		// Admins and librarians can see all borrows
		return borrowHandler.GetAllBorrows(c)
	})

	// Librarian and admin routes for borrow management
	borrowManagement := borrows.Group("", middleware.RoleRequired("admin", "librarian"))
	borrowManagement.Get("/all", borrowHandler.GetAllBorrows)
	borrowManagement.Get("/active", borrowHandler.GetActiveBorrows)
	borrowManagement.Get("/overdue", borrowHandler.GetOverdueBorrows)
	borrowManagement.Get("/:id", borrowHandler.GetBorrowByID)
	borrowManagement.Put("/:id", borrowHandler.UpdateBorrow)
	borrowManagement.Put("/:id/return", borrowHandler.ReturnBook)
	borrowManagement.Get("/user/:userId", borrowHandler.GetBorrowsByUser)
	borrowManagement.Get("/book/:bookId", borrowHandler.GetBorrowsByBook)
	borrowManagement.Get("/user/:userId/history", borrowHandler.GetBorrowHistory)

	// User-specific routes (users can access their own data)
	userSpecific := protected.Group("/my")
	userSpecific.Get("/borrows", func(c *fiber.Ctx) error {
		// Set user ID from token to params
		userID := c.Locals("user_id").(uint)
		c.Params("userId", string(rune(userID)))
		return borrowHandler.GetBorrowsByUser(c)
	})
	userSpecific.Get("/history", func(c *fiber.Ctx) error {
		// Set user ID from token to params
		userID := c.Locals("user_id").(uint)
		c.Params("userId", string(rune(userID)))
		return borrowHandler.GetBorrowHistory(c)
	})
}
