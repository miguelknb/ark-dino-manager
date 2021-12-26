package auth

import (
	"github.com/gin-gonic/gin"
)

// Routes for /auth
func Routes(r *gin.Engine) {
	auth := r.Group("/auth")

	{
		auth.POST("/login", login)
		auth.POST("/register", register)
	}
}
