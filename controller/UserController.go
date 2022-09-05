package controller

import (
	"efishery-project/entities"
	"efishery-project/middlewares"
	"efishery-project/service"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller_User struct {
	authservice middlewares.JWTService
	serviceuser service.ServiceUser
}

func NewControllerUser(authservice middlewares.JWTService, serviceuser service.ServiceUser) *Controller_User {
	return &Controller_User{authservice, serviceuser}
}

func (cu *Controller_User) CreateUserController(c echo.Context) error {
	var createInput entities.User

	err := c.Bind(&createInput)

	if err != nil {
		log.Panic("Failed to bind input")
		return c.JSON(http.StatusBadRequest, err)
	}

	createUser, errCreateUser := cu.serviceuser.CreateUserService(createInput)

	if errCreateUser != nil {
		log.Println("Failed to create user")
		return c.JSON(http.StatusBadRequest, errCreateUser.Error())
	}

	return c.JSON(http.StatusOK, createUser)
}

func (cu *Controller_User) LoginController(c echo.Context) error {
	var loginUser entities.User

	err := c.Bind(&loginUser)

	if err != nil {
		log.Panic("Failed to bind input")
		return c.JSON(http.StatusBadRequest, err)
	}

	login, errLogin := cu.serviceuser.LoginService(loginUser)

	if errLogin != nil {
		log.Println("Failed to login")
		return c.JSON(http.StatusBadRequest, errLogin.Error())
	}

	token, errToken := cu.authservice.GenerateToken(int(login.ID))

	if errToken != nil {
		log.Println(token)
		return c.String(http.StatusBadRequest, "Failed to create token")
	}

	c.Response().Header().Set("Authorization", token)

	return c.JSON(http.StatusOK, login)
}

func (cu *Controller_User) GetUserByIdController(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.String(http.StatusBadRequest, "Failed to convert id")
	}

	user, errUser := cu.serviceuser.GetUserByIdService(idInt)

	if errUser != nil {
		return c.String(http.StatusBadRequest, "Failed to get user")
	}

	return c.JSON(http.StatusOK, user)
}
