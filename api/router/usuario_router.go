package router

import (
	"github.com/labstack/echo/v4"
	"github.com/user0608/api_rest/api/paths"
	"github.com/user0608/api_rest/cmd/injectors"
)

func updradeUsuario(e *echo.Echo) {
	h := injectors.GetUsuarioHandler()
	g := e.Group(paths.USUARIO)
	g.POST("/login", h.LoginHandler)
	g.GET("/", h.FindAll)
}
