package users

import (
	"database/sql"
	"log"

	db "github.com/khwj/hackernews/internal/pkg/db/mysql"
	"golang.org/x/crypto/bcrypt"
)

// User is a struct that represent the users we get from database
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Create ...
func (user *User) Create() {
	stmt, err := db.DB.Prepare("INSERT INTO users(username, password) VALUES(?, ?)")
	if err != nil {
		log.Fatal("Error preparing insert user statement, ", err)
		return
	}

	hash, err := HashPassword(user.Password)
	if err != nil {
		log.Fatal("Error hashing password, ", err)
		return
	}

	_, err = stmt.Exec(user.Username, hash)
	if err != nil {
		log.Fatal("Error inserting user into database, ", err)
		return
	}
}

// GetUserIDByUsername check if a user exists in database by given username
func GetUserIDByUsername(username string) (int, error) {
	stmt, err := db.DB.Prepare("SELECT id FROM users WHERE username = ?")
	if err != nil {
		log.Fatal("Error preparing get user id statement, ", err)
		return -1, err
	}
	row := stmt.QueryRow(username)

	var id int
	err = row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
		}
		log.Println(err)
		return -1, err
	}
	return id, nil
}

// HashPassword hashes given password
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(hash), err
}

// VerifyPassword compares raw password with it's hashed values
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Authenticate ...
func (user *User) Authenticate() bool {
	stmt, err := db.DB.Prepare("SELECT password FROM users WHERE username = ?")
	if err != nil {
		log.Fatal("Error preparing authentication statement")
	}
	row := stmt.QueryRow(user.Username)

	var hash string
	err = row.Scan(&hash)
	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		log.Fatal(err)
	}

	return VerifyPassword(user.Password, hash)
}
