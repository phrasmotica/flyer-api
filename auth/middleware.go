package auth

import (
	"net/http"
	"phrasmotica/flyer-api/logging"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

// taken from https://stackoverflow.com/a/29439630
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func TokenAuth(optional bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		success, tokenString := parseToken(c)

		if !success {
			if optional {
				c.Next()
				return
			}

			logging.Error.Println("Request does not contain an access token")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, err := ValidateToken(tokenString)
		if err != nil {
			logging.Error.Println(err.Error())
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("token", tokenString)

		c.Set("username", claims.Username)
		c.Set("email", claims.Email)

		roles := claims.ResourceAccess[claims.Azp].Roles
		c.Set("roles", roles)

		logging.Info.Printf("Roles for user %s\n", claims.Username)
		logging.Info.Println(roles)

		c.Next()
	}
}

func CheckPermission(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.GetString("username")
		roles := c.GetStringSlice("roles")

		if !slices.Contains(roles, permission) {
			logging.Error.Printf("User %s does not have the %s permission\n", username, permission)
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}

func parseToken(c *gin.Context) (bool, string) {
	header := c.GetHeader("Authorization")
	splitToken := strings.Split(header, "Bearer ")

	if len(splitToken) != 2 {
		return false, ""
	}

	return true, splitToken[1]
}
