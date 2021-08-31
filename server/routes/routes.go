package routes

import (
	"github.com/gdgrosario/movieland/routes/auth"
	"github.com/gdgrosario/movieland/routes/posts"
	"github.com/gdgrosario/movieland/routes/users"

	"github.com/labstack/echo/v4"
)

// Routes : Init Routes
func Routes(g *echo.Group) {
	auth.AuthRouter{}.Init(g.Group("/auth"))
	posts.PostsRouter{}.Init(g.Group("/posts"))
  users.UsersRouter{}.Init(g.Group("/users"))
}
