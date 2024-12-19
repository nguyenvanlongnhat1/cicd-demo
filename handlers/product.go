package handlers

import (
	"cicd-demo/services"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService services.ProductService
}

func ProvideProductHandler(
	productService services.ProductService,
) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) GetProducts(c echo.Context) error {
	h.productService.GetProductDetail(c, 1)
	return nil
}
