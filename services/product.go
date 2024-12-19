package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductService interface {
	GetProductDetail(c echo.Context, id int) error
}

type productService struct {
}

func ProvideProductService() ProductService {
	return &productService{}
}

func (s *productService) GetProductDetail(c echo.Context, id int) error {
	return c.String(http.StatusOK, "Xin chào")
}
