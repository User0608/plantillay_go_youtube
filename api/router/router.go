package router

import "github.com/labstack/echo/v4"

func Updrade(e *echo.Echo) {
	updradeUsuario(e)
	updradeEmpleado(e)
}
