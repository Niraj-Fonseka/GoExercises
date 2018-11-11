package auth

import (
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func Authenticator(userId string, password string, c *gin.Context) (string, bool) {
	if (userId == "admin" && password == "admin") || (userId == "test" && password == "test") {
		return userId, true
	}

	return userId, false
}

func Authorizator(userId string, c *gin.Context) bool {
	if userId == "admin" {
		return true
	}

	return false
}

func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

//func InitAuthentication() *jwt.GinJWTMiddleware {
func InitAuthentication() *jwt.GinJWTMiddleware {
	// the jwt middleware

	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte("secret key"),
		Timeout:       10 * time.Hour,
		MaxRefresh:    time.Hour,
		Authenticator: Authenticator,
		Authorizator:  Authorizator,
		Unauthorized:  Unauthorized,
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "JWT",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}

	return authMiddleware
}
