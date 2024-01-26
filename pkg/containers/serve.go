package containers

import (
	"book-crud/pkg/config"
	"book-crud/pkg/connection"
	"book-crud/pkg/controllers"
	"book-crud/pkg/repositories"
	"book-crud/pkg/routes"
	"book-crud/pkg/services"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func Serve(e *echo.Echo) {

	//config initialization
	config.SetConfig()

	//database initializations
	db := connection.GetDB()

	// repository initialization
	bookRepo := repositories.BookDBInstance(db)
	userRepo := repositories.UserDBInstance(db)

	//service initialization
	bookService := services.BookServiceInstance(bookRepo)
	userService := services.UserServiceInstance(userRepo)

	//controller initialization
	bookCtr := controllers.NewBookController(bookService)
	userCtr := controllers.NewUserController(userService)

	//route initialization
	bookRoute := routes.BookRoutes(e, bookCtr)
	userRoute := routes.UserRoutes(e, userCtr)

	bookRoute.InitBookRoute()
	userRoute.InitUserRoute()
	// starting server
	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))

}
