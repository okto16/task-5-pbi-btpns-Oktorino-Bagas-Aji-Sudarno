package middlewares

import (
    "github.com/gin-gonic/gin"
    "golang-api/helpers"
    "net/http"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        userID, err := helpers.VerifyToken(tokenString)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }
        c.Set("userID", userID)
        c.Next()
    }
}
