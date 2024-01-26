package controllers

import (
	"book-crud/pkg/config"
	"book-crud/pkg/domain"
	"book-crud/pkg/models"
	"book-crud/pkg/types"
	"book-crud/pkg/utils"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type IUserController interface {
	CreateUser(e echo.Context) error
	GetUsers(e echo.Context) error
	GetUsersByUsername(e echo.Context) error
	UpdateUser(e echo.Context) error
	DeleteUser(e echo.Context) error
}

// to access the methods of service and repo
type UserController struct {
	userSvc domain.IUserService
}

func NewUserController(userSvc domain.IUserService) UserController {
	return UserController{
		userSvc: userSvc,
	}
}

func (uc *UserController) CreateUser(e echo.Context) error {
	reqUser := &types.UserRequest{}
	if err := e.Bind(reqUser); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqUser.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	user := &models.UserDetail{
		Username: reqUser.Username,
		Password: utils.HashPassword(reqUser.Password),
	}

	if err := uc.userSvc.CreateUser(user); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, "UserDetail was created successfully")
}
func (uc *UserController) GetUsers(e echo.Context) error {
	tempUserID := e.QueryParam("userID")
	userID, err := strconv.ParseInt(tempUserID, 0, 0)
	if err != nil && tempUserID != "" {
		return e.JSON(http.StatusBadRequest, "Enter a valid user ID")
	}
	user, err := uc.userSvc.GetUsers(uint(userID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, user)
}

func (uc *UserController) GetUsersByUsername(e echo.Context) error {
	username := e.QueryParam("username")
	user, err := uc.userSvc.GetUsersByUsername(username)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUser(e echo.Context) error {
	reqUser := &types.UserRequest{}
	if err := e.Bind(reqUser); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqUser.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	tempUserID := e.Param("userID")
	userID, err := strconv.ParseInt(tempUserID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Enter a valid user ID")
	}
	existingUser, err := uc.userSvc.GetUsers(uint(userID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	_ = existingUser
	updatedUser := &models.UserDetail{
		ID:       uint(userID),
		Username: reqUser.Username,
		Password: utils.HashPassword(reqUser.Password),
	}

	if err := uc.userSvc.UpdateUser(updatedUser); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, "UserDetail was updated successfully")
}

func (uc *UserController) DeleteUser(e echo.Context) error {
	tempUserID := e.Param("userID")
	userID, err := strconv.ParseInt(tempUserID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	_, err = uc.userSvc.GetUsers(uint(userID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	if err := uc.userSvc.DeleteUser(uint(userID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "UserDetail was deleted successfully")
}

func (uc *UserController) Login(c echo.Context) error {
	config := config.LocalConfig
	var auth types.AuthRequest
	if err := c.Bind(&auth); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid Data")
	}

	// Get the user from the database by username and password
	users, err := uc.userSvc.GetUsersByUsername(auth.Username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user := users[0]
	if err := utils.ComparePassword(user.Password, auth.Password); err == nil {
		now := time.Now().UTC()
		ttl := time.Minute * 1
		claims := jwt.StandardClaims{
			Subject:   auth.Username,
			ExpiresAt: now.Add(ttl).Unix(),
			IssuedAt:  now.Unix(),
			NotBefore: now.Unix(),
		}
		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.JwtSecret))
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, "Invalid token")
		}
		return c.JSON(http.StatusOK, token)
	} else {
		return c.JSON(http.StatusUnauthorized, "Login Failed")
	}
}
