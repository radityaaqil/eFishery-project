package middlewares

import (
	"efishery-project/service"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(authService JWTService, userService service.ServiceUser, next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			fmt.Println("token", err)
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("claims", ok)
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}

		userID := int(claims["id"].(float64))
		user, err := userService.GetUserByIdService(userID)
		if err != nil {
			fmt.Println("userID", err)
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}
		c.Set("currentUser", user)
		return next(c)
	}
}
