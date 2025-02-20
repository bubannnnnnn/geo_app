package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates the JWT token in the Authorization header.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the JWT token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			return
		}

		// Extract the token from the "Bearer <token>" format
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		if tokenString == authHeader { // No "Bearer " prefix
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}

		// Validate the JWT token (replace with your actual validation logic)
		// Here, we'll just check if the token is not empty for demonstration purposes
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT token"})
			return
		}

		// TODO: Verify JWT signature and claims
		// You would typically use a library like "github.com/golang-jwt/jwt/v5" for this

		// For demonstration purposes, let's assume the token is valid
		// You might want to extract user information from the token and store it in the context:
		// userID, err := extractUserIDFromToken(tokenString)
		// if err != nil {
		//    c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT token"})
		//    return
		// }
		// c.Set("userID", userID)

		c.Next() // Proceed to the next handler
	}
}
