package routes

import (
	"myapp/config"
	"myapp/controllers"
	"myapp/middlewares"
	validation "myapp/validations"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var envVars = config.GetEnv()

func User(e *echo.Echo) {
	// Middleware login
	isLogin := middleware.JWT([]byte(envVars.Jwt.SecretKey))
	users := e.Group("/users", isLogin, middlewares.CheckAdminRole)
	users.POST("", controllers.CreateUser, validation.UserCreateBody) // --> ok
	users.GET("", controllers.GetListUser)
	users.PUT("/:id", controllers.UpdateUser, validation.ValidateID, validation.UserUpdateBody) // --> ok
	users.DELETE("/:id", controllers.DeleteUser, validation.ValidateID)
}
