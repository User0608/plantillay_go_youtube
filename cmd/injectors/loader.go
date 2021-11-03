package injectors

import (
	"sync"

	"github.com/user0608/api_rest/configs"
	"github.com/user0608/api_rest/database"
	"github.com/user0608/api_rest/handlers"
	"github.com/user0608/api_rest/repos"
	"github.com/user0608/api_rest/services"
	"gorm.io/gorm"
)

var ones sync.Once

func LoadConfig(configFilePath string) error {
	var errr error
	ones.Do(func() {
		if conf, err := configs.LoadDBConfigs(configFilePath); err != nil {
			errr = err
			return
		} else {
			con, er := database.GetDBConnextion(conf)
			if err != nil {
				errr = er
				return
			}
			initRepository(con)
			initServices()
			initHandlers()
		}
	})
	return errr
}
func initRepository(connextion *gorm.DB) {
	usuarioRepository = repos.NewUsuarioRepository(connextion)
}
func initServices() {
	usuarioService = services.NewUsuarioService(usuarioRepository)
}
func initHandlers() {
	usuarioHandler = handlers.NewUsuarioHandler(usuarioService)
}
