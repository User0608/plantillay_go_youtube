package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func OKResponse(c echo.Context, payload interface{}) error {
	return c.JSON(http.StatusOK, payload)
}
func BadReponse(c echo.Context, payload interface{}) error {
	return nil
}
