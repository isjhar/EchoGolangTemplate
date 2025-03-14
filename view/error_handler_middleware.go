// elapsed_time_middleware.go
package view

import (
	"errors"
	"isjhar/template/echo-golang/domain/entities"
	"isjhar/template/echo-golang/utils"
	"isjhar/template/echo-golang/view/dto"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	// Check if the context has been sent already
	if c.Response().Committed {
		return
	}

	if err != nil {
		statusCode := http.StatusInternalServerError
		msg := "Internal Server Error"

		if he, ok := err.(*echo.HTTPError); ok {
			statusCode = he.Code
			msg = he.Message.(string)
		} else if _, ok := err.(validator.ValidationErrors); ok {
			utils.LogError("invalid params at [%s] url %s : %v", c.Request().Method, c.Request().URL.Path, err)
			statusCode = http.StatusBadRequest
			msg = entities.InvalidParams.Error()
		} else if se, ok := err.(*entities.ServiceError); ok {
			if !errors.Is(err, entities.InternalServerError) {
				if errors.Is(err, entities.Forbidden) {
					statusCode = http.StatusForbidden
				} else {
					statusCode = http.StatusBadRequest
				}
			}
			msg = se.Error()
		}

		c.JSON(statusCode, dto.ApiResponse{
			Message: msg,
		})
	}
}
