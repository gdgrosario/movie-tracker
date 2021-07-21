package main

import (
	"github.com/gdgrosario/movieland/database"
	"github.com/gdgrosario/movieland/lib"
	"github.com/gdgrosario/movieland/lib/middlewares"
	"github.com/gdgrosario/movieland/routes"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	godotenv.Load(".env.development")
}

func main() {
	db := database.ConnectDB()
	defer db.Close()

	e := echo.New()

	e.Validator = &lib.CustomValidator{Validator: validator.New()}

	e.Use(middlewares.ContextDB(db))

	routes.Routes(e.Group(""))

	e.Logger.Fatal(e.Start(":3000"))
}
