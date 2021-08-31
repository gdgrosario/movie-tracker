package auth

import (
	"net/http"
	"time"

	"github.com/gdgrosario/movieland/database/models"
	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// Register : Register Router
func (AuthRouter) Register(c echo.Context) error {
	type RequestBody struct {
		Username       string `json:"username" validate:"required"`
		Password       string `json:"password" validate:"required"`
    Email          string `json:"email" validate:"required"`
    DisplayName    string `json:"displayName" validate:"required"`
	}

	var body RequestBody

	if err := c.Bind(&body); err != nil {
		return err
	}
	if err := c.Validate(&body); err != nil {
		return err
	}

	db, _ := c.Get("db").(*gorm.DB)

	user := models.User{
		Username:     body.Username,
		PasswordHash: body.Password,
    Email:        body.Email,
		DisplayName:  body.DisplayName,
	}

	user.HashPassword()
	db.Create(&user)

	token, _ := user.GenerateToken()

  return c.JSON(http.StatusOK, echo.Map{
		"token": token,
		"user":  user,
	})
}

// Login : Login Router
func (AuthRouter) Login(c echo.Context) error {
	type RequestBody struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	var body RequestBody

	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	if err := c.Validate(&body); err != nil {
		return err
	}

	db, _ := c.Get("db").(*gorm.DB)

	var user models.User

	if err := db.Where("email = ?", body.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusConflict, "Invalid credentials")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Password)) != nil {
		return c.JSON(http.StatusConflict, "Invalid credentials")
	}

	token, _ := user.GenerateToken()

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
		"user":  user,
	})
}

// Logout : Logout Router
func (AuthRouter) Logout(c echo.Context) error {
	tokenCookie, _ := c.Get("tokenCookie").(*http.Cookie)

	tokenCookie.Value = ""
	tokenCookie.Expires = time.Unix(0, 0)

	c.SetCookie(tokenCookie)

	return c.NoContent(200)
}
