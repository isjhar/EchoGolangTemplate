package controllers

import (
	"isjhar/template/echo-golang/view"
	"isjhar/template/echo-golang/view/dto"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := &dto.LoginParams{}

		if err := c.Bind(data); err != nil {
			return c.JSON(http.StatusBadRequest, dto.ApiResponse{
				Message: err.Error(),
			})
		}

		if err := c.Validate(data); err != nil {
			return c.JSON(http.StatusBadRequest, dto.ApiResponse{
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, dto.ApiResponse{})
	}
}

func IsAuth() echo.HandlerFunc {
	return func(c echo.Context) error {
		authorizedContext := c.(*view.AuthorizedContext)
		log.Println(authorizedContext.User)
		return c.JSON(http.StatusOK, dto.ApiResponse{})
	}
}
