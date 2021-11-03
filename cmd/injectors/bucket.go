package injectors

import (
	"github.com/user0608/api_rest/handlers"
	"github.com/user0608/api_rest/repos"
	"github.com/user0608/api_rest/services"
	"gorm.io/gorm"
)

var (
	Connextion *gorm.DB
	//repos
	usuarioRepository *repos.UsuarioRepository

	//service
	usuarioService *services.UsuarioService

	//handlers
	usuarioHandler *handlers.UsuarioHandler
)
