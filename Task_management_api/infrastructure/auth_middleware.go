package infrastructure

import (
	"net/http"
	"strings"
	domain "task_manager/domain"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	JwtService domain.JwtSvc
}

func NewAuthMiddleware(jwtService domain.JwtSvc) *AuthMiddleware {
	return &AuthMiddleware{
		JwtService: jwtService,
	}
}

func (m *AuthMiddleware) AuthorizeJWT(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		tokenString = strings.TrimSpace(tokenString)

		token, err := m.JwtService.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse token claims"})
			return
		}

		userID, ok := claims["user_id"].(string)
		userType, ok2 := claims["user_type"].(string)
		if !ok || !ok2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		if len(allowedRoles) > 0 {
			roleAllowed := false
			for _, role := range allowedRoles {
				if strings.ToLower(role) == userType {
					roleAllowed = true
					break
				}
			}
			if !roleAllowed {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
				return
			}
		}

		c.Set("user_id", userID)
		c.Set("user_type", userType)

		c.Next()
	}
}
