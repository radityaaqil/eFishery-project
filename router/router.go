package router

import (
	"efishery-project/controller"
	"efishery-project/middlewares"
	"efishery-project/repository"
	"efishery-project/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *echo.Echo {
	//new echo
	e := echo.New()

	//Middleware
	AuthService := middlewares.NewJwtService()

	//User
	UserRepository := repository.NewRepositoryUser(db)
	UserService := service.NewServiceUser(UserRepository)
	UserController := controller.NewControllerUser(AuthService, UserService)

	//User Router
	e.POST("/signup", UserController.CreateUserController)
	e.POST("/signin", UserController.LoginController)
	e.GET("/user/:id", middlewares.AuthMiddleware(AuthService, UserService, UserController.GetUserByIdController))

	return e
}
