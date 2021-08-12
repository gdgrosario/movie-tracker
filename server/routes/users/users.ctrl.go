package users

import (
	"net/http"

	"github.com/gdgrosario/movieland/database/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func (UsersRouter) UsersGetAll(c echo.Context) error {
  db, _ := c.Get("db").(*gorm.DB)

  var users models.User

  result := db.Find(&users)

  return c.JSON(http.StatusOK, result)
}


