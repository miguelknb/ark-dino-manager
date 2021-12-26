package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/postgres"
	"github.com/gin-gonic/gin"

	"github.com/miguelknb/ark-dino-manager/db"
	"github.com/miguelknb/ark-dino-manager/endpoints/auth"
)

func main() {
	r := gin.Default()

	// setup pool
	db.Init()

	// setup cookie session
	sess, err := postgres.NewStore(db.DB, []byte(os.Getenv("SESSION_SECRET")))
	if err != nil {
		log.Fatalf("Coult not initialize cookie store: %v", err)
	}

	sess.Options(sessions.Options{
		MaxAge:   30 * 60 * 60 * 24,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
	})
	r.Use(sessions.Sessions("login", sess))

	// authentication routes
	auth.Routes(r)

	// endpoint GET /
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Everything is ok."})
	})

	// run server
	r.Run()
}
