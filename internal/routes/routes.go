package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yooerizkilab/library-system/internal/database"
	"github.com/yooerizkilab/library-system/internal/handlers"
	"github.com/yooerizkilab/library-system/internal/repositories"
	"github.com/yooerizkilab/library-system/internal/services"
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

	// User routes
	users := v1.Group("/users")
	users.Post("/", userHandler.CreateUser)
	users.Get("/", userHandler.GetAllUsers)
	users.Get("/search", userHandler.SearchUsers)
	users.Get("/:id", userHandler.GetUserByID)
	users.Put("/:id", userHandler.UpdateUser)
	users.Delete("/:id", userHandler.DeleteUser)

	// Book routes
	books := v1.Group("/books")
	books.Post("/", bookHandler.CreateBook)
	books.Get("/", bookHandler.GetAllBooks)
	books.Get("/search", bookHandler.SearchBooks)
	books.Get("/available", bookHandler.GetAvailableBooks)
	books.Get("/category/:category", bookHandler.GetBooksByCategory)
	books.Get("/:id", bookHandler.GetBookByID)
	books.Put("/:id", bookHandler.UpdateBook)
	books.Delete("/:id", bookHandler.DeleteBook)

	// Borrow routes
	borrows := v1.Group("/borrows")
	borrows.Post("/", borrowHandler.BorrowBook)
	borrows.Get("/", borrowHandler.GetAllBorrows)
	borrows.Get("/active", borrowHandler.GetActiveBorrows)
	borrows.Get("/overdue", borrowHandler.GetOverdueBorrows)
	borrows.Get("/:id", borrowHandler.GetBorrowByID)
	borrows.Put("/:id", borrowHandler.UpdateBorrow)
	borrows.Put("/:id/return", borrowHandler.ReturnBook)
	borrows.Get("/user/:userId", borrowHandler.GetBorrowsByUser)
	borrows.Get("/book/:bookId", borrowHandler.GetBorrowsByBook)
	borrows.Get("/user/:userId/history", borrowHandler.GetBorrowHistory)
}
