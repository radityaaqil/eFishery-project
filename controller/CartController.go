package controller

import (
	"efishery-project/entities"
	"efishery-project/service"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller_Cart struct {
	servicecart service.ServiceCart
}

func NewControllerCart(servicecart service.ServiceCart) *Controller_Cart {
	return &Controller_Cart{servicecart}
}

func (co *Controller_Cart) InsertToCartController(c echo.Context) error {
	var createInput entities.Cart

	err := c.Bind(&createInput)

	if err != nil {
		log.Panic("Failed to bind input")
		return c.JSON(http.StatusBadRequest, err)
	}

	user_id_tkn := c.Get("currentUser").(entities.User)
	createInput.User_id = int(user_id_tkn.ID)

	createInputCart, errCreateInputCart := co.servicecart.InsertToCartService(createInput)

	if errCreateInputCart != nil {
		log.Println("Failed to create cart")
		return c.JSON(http.StatusBadRequest, errCreateInputCart.Error())
	}

	return c.JSON(http.StatusOK, createInputCart)
}

func (co *Controller_Cart) GetCartController(c echo.Context) error {
	var user entities.Cart

	user_id_tkn := c.Get("currentUser").(entities.User)
	user.User_id = int(user_id_tkn.ID)

	getCartResult, errgetCartResult := co.servicecart.GetCartService(user.User_id)

	if errgetCartResult != nil {
		return c.String(http.StatusBadRequest, "Failed to load cart")
	}

	return c.JSON(http.StatusOK, getCartResult)
}
