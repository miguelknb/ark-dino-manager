package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/miguelknb/ark-dino-manager/db"
	"github.com/miguelknb/ark-dino-manager/endpoints/auth"
)

func main() {
	r := gin.Default()

	// setup pool
	db.Init()

	// setup cookie session
	r.Use(sessions.Sessions("login-session", sessions.NewCookieStore([]byte("TODOsecret"))))

	// authentication routes
	auth.Routes(r)

	// endpoint GET /
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Everything is ok."})
	})

	// run server
	r.Run()
}
