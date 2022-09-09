package controller

import (
	"efishery-project/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller_Product struct {
	serviceproduct service.ServiceProduct
}

func NewControllerProduct(serviceproduct service.ServiceProduct) *Controller_Product {
	return &Controller_Product{serviceproduct}
}

func (cu *Controller_Product) GetListProductController(c echo.Context) error {
	page := c.QueryParam("page")
	pageInt, _ := strconv.Atoi(page)
	category := c.QueryParam("category")
	minPrice := c.QueryParam("minPrice")
	minPriceInt, _ := strconv.Atoi(minPrice)
	maxPrice := c.QueryParam("maxPrice")
	maxPriceInt, _ := strconv.Atoi(maxPrice)

	productList, errProductList := cu.serviceproduct.GetListProductService(pageInt, category, minPriceInt, maxPriceInt)

	if errProductList != nil {
		return c.String(http.StatusBadRequest, "Failed to retrieve products")
	}

	if len(productList) == 0 {
		return c.String(http.StatusOK, "Sorry, products unavailable")
	}

	return c.JSON(http.StatusOK, productList)
}

func (cu *Controller_Product) GetProductDetailController(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.String(http.StatusBadRequest, "Failed to convert id")
	}

	productDetail, errProductDetail := cu.serviceproduct.GetProductDetailService(idInt)

	if errProductDetail != nil {
		return c.String(http.StatusBadRequest, "Failed to retrieve products")
	}

	return c.JSON(http.StatusOK, productDetail)
}
