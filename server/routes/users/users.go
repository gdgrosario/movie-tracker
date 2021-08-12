package users

import "github.com/labstack/echo/v4"

// Struct for Users Router

type UsersRouter struct {}

func (ctrl UsersRouter) Init(g *echo.Group) {
  g.GET("/all", ctrl.UsersGetAll)
}

