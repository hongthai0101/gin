package middlewares

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hongthai0101/golang-gin/types"
	"github.com/hongthai0101/golang-gin/utils"
	"net/http"
	"strings"
)

//Authentication is for auth middleware
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if len(authHeader) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authentication header is missing",
			})
			return
		}
		temp := strings.Split(authHeader, "Bearer")
		if len(temp) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token not Bearer"})
			return
		}
		tokenString := strings.TrimSpace(temp[1])
		tokenSlice := strings.Split(tokenString, ".")
		tokenData := utils.FindSliceByKey(tokenSlice, 1, "")

		if data, err := base64.StdEncoding.DecodeString(tokenData); err == nil {
			tokenPayload := types.TokenPayload{}
			errUnmarshal := json.Unmarshal(data, &tokenPayload)
			if errUnmarshal != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "Token is not valid Unmarshal",
				})
				return
			}
			c.Set("user", tokenPayload)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token is not valid",
			})
			return
		}
	}
}

//ErrorHandler is for global error
func ErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": c.Errors,
		})
	}
}