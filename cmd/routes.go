package cmd

import (
	"cicd-demo/handlers"

	"github.com/labstack/echo/v4"
)

func Router(
	handlers *handlers.ProductHandler,
) *echo.Echo {
	e := echo.New()

	// e.GET("/products/:id", GetProductDetail)
	e.GET("/", handlers.GetProducts)

	return e
}
