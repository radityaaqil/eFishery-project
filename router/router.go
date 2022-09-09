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

	//Products
	ProductRepository := repository.NewRepositoryProduct(db)
	ProductService := service.NewServiceProduct(ProductRepository)
	ProductController := controller.NewControllerProduct(ProductService)

	//Product Router
	e.GET("/products", ProductController.GetListProductController)
	e.GET("/product/:id", ProductController.GetProductDetailController)

	//Cart
	CartRepository := repository.NewRepositoryCart(db)
	CartService := service.NewServiceCart(CartRepository)
	CartController := controller.NewControllerCart(CartService)

	//Cart Router
	e.POST("/cart", middlewares.AuthMiddleware(AuthService, UserService, CartController.InsertToCartController))
	e.GET("/cart", middlewares.AuthMiddleware(AuthService, UserService, CartController.GetCartController))

	//Transaction
	TransactionRepository := repository.NewRepositoryTransaction(db)
	TransactionService := service.NewServiceTransaction(TransactionRepository)
	TransactionController := controller.NewControllerTransaction(TransactionService)

	//Transaction Router
	e.POST("transaction", middlewares.AuthMiddleware(AuthService, UserService, TransactionController.CreateTransactionController))
	e.POST("transaction/paymentslip/:transaction_id", middlewares.AuthMiddleware(AuthService, UserService, TransactionController.InsertPaymentSlipController))
	e.GET("transactions", middlewares.AuthMiddleware(AuthService, UserService, TransactionController.GetUserTransactionsController))

	return e
}
