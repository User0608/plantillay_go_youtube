package injectors

import "github.com/user0608/api_rest/handlers"

func GetUsuarioHandler() *handlers.UsuarioHandler {
	return usuarioHandler
}
