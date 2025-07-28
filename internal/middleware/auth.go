package middleware

import (
	"strings"
	
	"github.com/yooerizkilab/library-system/pkg/response"
	"github.com/yooerizkilab/library-system/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// AuthRequired middleware untuk memverifikasi JWT token
func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get token from Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return response.BadRequest(c, "Authorization header required", nil)
		}

		// Check if header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return response.BadRequest(c, "Invalid authorization header format", nil)
		}

		// Extract token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			return response.BadRequest(c, "Token required", nil)
		}

		// Validate token
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			return response.BadRequest(c, "Invalid or expired token", err.Error())
		}

		// Store user info in context
		c.Locals("user_id", claims.UserID)
		c.Locals("user_email", claims.Email)
		c.Locals("user_role", claims.Role)

		return c.Next()
	}
}

// RoleRequired middleware untuk memverifikasi role user
func RoleRequired(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole := c.Locals("user_role").(string)

		for _, role := range allowedRoles {
			if userRole == role {
				return c.Next()
			}
		}

		return response.BadRequest(c, "Insufficient permissions", "Access denied for role: "+userRole)
	}
}

// Optional auth middleware - tidak wajib login tapi kalau ada token akan divalidasi
func OptionalAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString != "" {
				claims, err := utils.ValidateToken(tokenString)
				if err == nil {
					c.Locals("user_id", claims.UserID)
					c.Locals("user_email", claims.Email)
					c.Locals("user_role", claims.Role)
				}
			}
		}
		return c.Next()
	}
}