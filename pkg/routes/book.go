package routes

import (
	"book-crud/pkg/controllers"
	"book-crud/pkg/middlewares"

	"github.com/labstack/echo/v4"
)

type bookRoutes struct {
	echo    *echo.Echo
	bookCtr controllers.BookController
}

func BookRoutes(echo *echo.Echo, bookCtr controllers.BookController) *bookRoutes {
	return &bookRoutes{
		echo:    echo,
		bookCtr: bookCtr,
	}
}

func (bc *bookRoutes) InitBookRoute() {
	e := bc.echo
	bc.initBookRoutes(e)
}

func (bc *bookRoutes) initBookRoutes(e *echo.Echo) {
	//grouping route endpoints
	book := e.Group("/bookstore")

	book.GET("/book", bc.bookCtr.GetBook)

	// Apply middleware
	book.Use(middlewares.Auth)

	book.POST("/book", bc.bookCtr.CreateBook)
	book.PUT("/book/:bookID", bc.bookCtr.UpdateBook)
	book.DELETE("/book/:bookID", bc.bookCtr.DeleteBook)
}
