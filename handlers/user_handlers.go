package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"example/web-service-gin/db"
	"example/web-service-gin/models"
)



func GetUsers(c echo.Context) error {
	rows, err := db.DB.Query("SELECT id, name FROM users")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}
	defer rows.Close()

	users := make([]models.User, 0)

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		}
		users = append(users, user)
	}

	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")

	rows, err := db.DB.Query("SELECT id, name FROM users where id = $1", id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}
	defer rows.Close()

	// Check if there are any rows
	if !rows.Next() {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}
	fmt.Print(rows)

	// Scan the values from the row
	var user models.User
	err = rows.Scan(&user.ID, &user.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	// Return the user data as JSON
	return c.JSON(http.StatusOK, user)
}


func CreateUser(c echo.Context) error {
	var newUser models.User
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