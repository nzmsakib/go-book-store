package routes

import (
	"book-crud/pkg/controllers"
	"book-crud/pkg/middlewares"

	"github.com/labstack/echo/v4"
)

type userRoutes struct {
	echo    *echo.Echo
	userCtr controllers.UserController
}

func UserRoutes(echo *echo.Echo, userCtr controllers.UserController) *userRoutes {
	return &userRoutes{
		echo:    echo,
		userCtr: userCtr,
	}
}

func (uc *userRoutes) InitUserRoute() {
	e := uc.echo
	uc.initUserRoutes(e)
}

func (uc *userRoutes) initUserRoutes(e *echo.Echo) {
	//grouping route endpoints
	user := e.Group("/bookstore")

	user.GET("/user/:userID", uc.userCtr.GetUsers)
	//login
	user.POST("/auth/login", uc.userCtr.Login)
	// register
	user.POST("/user", uc.userCtr.CreateUser)

	// Apply middleware
	user.Use(middlewares.Auth)

	user.PUT("/user/:userID", uc.userCtr.UpdateUser)
	user.DELETE("/user/:userID", uc.userCtr.DeleteUser)
}
