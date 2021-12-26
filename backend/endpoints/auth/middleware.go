package auth

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Authentication Required middleware
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	userid := session.Get("userid")

	if userid == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	c.Next()
}
