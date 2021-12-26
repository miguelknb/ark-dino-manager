package db

// This is the type for the "users" table
type User struct {
	ID            int64
	Username      string
	Email         string
	Password_hash []byte
}
