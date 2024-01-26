package controllers

import (
	"book-crud/pkg/domain"
	"book-crud/pkg/models"
	"book-crud/pkg/types"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type IBookController interface {
	CreateBook(e echo.Context) error
	GetBook(e echo.Context) error
	UpdateBook(e echo.Context) error
	DeleteBook(e echo.Context) error
}

// to access the methods of service and repo
type BookController struct {
	bookSvc domain.IBookService
}

func NewBookController(bookSvc domain.IBookService) BookController {
	return BookController{
		bookSvc: bookSvc,
	}
}

func (bs *BookController) CreateBook(e echo.Context) error {
	reqBook := &types.BookRequest{}
	if err := e.Bind(reqBook); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqBook.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	book := &models.BookDetail{
		BookName:    reqBook.BookName,
		Author:      reqBook.Author,
		Publication: reqBook.Publication,
	}

	if err := bs.bookSvc.CreateBook(book); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, "BookDetail was created successfully")
}
func (bs *BookController) GetBook(e echo.Context) error {
	tempBookID := e.QueryParam("bookID")
	bookID, err := strconv.ParseInt(tempBookID, 0, 0)
	if err != nil && tempBookID != "" {
		return e.JSON(http.StatusBadRequest, "Enter a valid book ID")
	}
	book, err := bs.bookSvc.GetBooks(uint(bookID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, book)
}

func (bs *BookController) UpdateBook(e echo.Context) error {
	reqBook := &types.BookRequest{}
	if err := e.Bind(reqBook); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqBook.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	tempBookID := e.Param("bookID")
	bookID, err := strconv.ParseInt(tempBookID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Enter a valid book ID")
	}
	existingBook, err := bs.bookSvc.GetBooks(uint(bookID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	updatedBook := &models.BookDetail{
		ID:          uint(bookID),
		BookName:    reqBook.BookName,
		Author:      reqBook.Author,
		Publication: reqBook.Publication,
	}
	if updatedBook.BookName == "" {
		updatedBook.BookName = existingBook[0].BookName
	}
	if updatedBook.Author == "" {
		updatedBook.Author = existingBook[0].Author
	}
	if updatedBook.Publication == "" {
		updatedBook.Publication = existingBook[0].Publication
	}
	if err := bs.bookSvc.UpdateBook(updatedBook); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, "BookDetail was updated successfully")
}

func (bs *BookController) DeleteBook(e echo.Context) error {
	tempBookID := e.Param("bookID")
	bookID, err := strconv.ParseInt(tempBookID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	_, err = bs.bookSvc.GetBooks(uint(bookID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	if err := bs.bookSvc.DeleteBook(uint(bookID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "BookDetail was deleted successfully")
}

func (bs *BookController) Login(c echo.Context) error {
	UserName := "admin"
	Password := "admin"

	type Auth struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	var auth Auth
	if err := c.Bind(&auth); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid Data")
	}

	if UserName == auth.UserName && Password == auth.Password {
		now := time.Now().UTC()
		ttl := time.Minute * 1
		claims := jwt.StandardClaims{
			Subject:   auth.UserName,
			ExpiresAt: now.Add(ttl).Unix(),
			IssuedAt:  now.Unix(),
			NotBefore: now.Unix(),
		}
		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("my-secret-key"))
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, "Invalid token")
		}
		return c.JSON(http.StatusOK, token)
	} else {
		return c.JSON(http.StatusUnauthorized, "Login Failed")
	}
}