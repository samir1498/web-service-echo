package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"example/web-service-gin/db"
)

// User represents a user in the database.
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetUsers handles the GET request to retrieve a list of users.
func GetUsers(c echo.Context) error {
	// Use db.DB to perform database operations
	rows, err := db.DB.Query("SELECT id, name FROM users")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}
	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		}
		users = append(users, user)
	}

	return c.JSON(http.StatusOK, users)
}

// CreateUser handles the POST request to create a new user.
func CreateUser(c echo.Context) error {
	// Use db.DB to perform database operations
	var newUser User
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	// Execute the INSERT statement with the RETURNING clause
	err := db.DB.QueryRow("INSERT INTO users (name) VALUES ($1) RETURNING id", newUser.Name).Scan(&newUser.ID)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	return c.JSON(http.StatusCreated, newUser)
}


func Hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}