package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/georgysavva/scany/sqlscan"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"golang.org/x/crypto/bcrypt"

	"github.com/miguelknb/ark-dino-manager/db"
)

// endpoint /auth/register POST JSON body
type registerBody struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// endpoint /auth/register
func register(c *gin.Context) {
	var body registerBody

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing fields"})
		return
	}

	// sanitize input
	username := strings.Trim(body.Username, " ")
	email := strings.Trim(body.Email, " ")
	password := strings.Trim(body.Password, " ")

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while processing your request"})
		return
	}

	// generate snowflake
	id := db.GenerateId()

	_, err = db.DB.Exec(
		"INSERT INTO users (id, username, email, password_hash) VALUES ($1, $2, $3, $4)",
		id, username, email, hash,
	)

	if err != nil {
		// postgres engine errors
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				c.JSON(http.StatusBadRequest, gin.H{"error": "An account already exists with this username or email"})
				return
			}
		}

		// other errors
		log.Printf("ERROR: %s\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while processing your request"})
		return
	}

	log.Printf("User %d created\n", id)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Account created with ID %d", id)})
}

// endpoint /auth/login POST JSON body
type loginBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// endpoint /auth/login
func login(c *gin.Context) {
	session := sessions.Default(c)

	var body loginBody
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing fields"})
		return
	}

	// sanitize input
	username := strings.Trim(body.Username, " ")
	password := strings.Trim(body.Password, " ")

	var users []*db.User
	err := sqlscan.Select(context.Background(), db.DB, &users, "SELECT id, username, email, password_hash FROM users WHERE email = $1 OR username = $1", username)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while processing your request"})
		return
	}

	if len(users) != 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect username or password"})
		return
	}

	user := users[0]

	if err := bcrypt.CompareHashAndPassword(user.Password_hash, []byte(password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect username or password"})
		return
	}

	// password is correct
	session.Set("userid", user.ID)
	if err := session.Save(); err != nil {
		log.Printf("ERROR: %s\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while processing your request"})
		return
	}

	log.Printf("User %d logged in\n", user.ID)
	c.JSON(http.StatusOK, gin.H{"message": "Authentication successful"})
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	userid := session.Get("userid")

	if userid == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User is not logged in"})
		return
	}

	session.Delete("userid")
	if err := session.Save(); err != nil {
		log.Printf("ERROR: %s\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while processing your request"})
		return
	}

	log.Printf("User %d logged out\n", userid)
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
