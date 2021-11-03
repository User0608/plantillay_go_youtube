package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/user0608/api_rest/errores"
	"github.com/user0608/api_rest/models"
	"github.com/user0608/api_rest/services"
	"github.com/user0608/api_rest/utils"
)

type UsuarioHandler struct {
	binder echo.DefaultBinder
	serv   *services.UsuarioService
}

func NewUsuarioHandler(s *services.UsuarioService) *UsuarioHandler {
	return &UsuarioHandler{serv: s, binder: echo.DefaultBinder{}}
}

func (h *UsuarioHandler) LoginHandler(c echo.Context) error {
	request := models.RequestSession{}
	if err := h.binder.BindBody(c, &request); err != nil {
		return errores.GenErrResponseEcho(c, err)
	}
	usuario, err := h.serv.Loggin(request)
	if err != nil {
		return errores.GenErrResponseEcho(c, err)
	}
	return utils.OKResponse(c, usuario)
}
func (h *UsuarioHandler) FindAll(c echo.Context) error {
	return utils.OKResponse(c, "Todo fue ok")
}
