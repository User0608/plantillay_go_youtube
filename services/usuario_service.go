package services

import (
	"errors"
	"strings"

	"github.com/user0608/api_rest/models"
	"github.com/user0608/api_rest/repos"
)

type UsuarioService struct {
	repository *repos.UsuarioRepository
}

func NewUsuarioService(r *repos.UsuarioRepository) *UsuarioService {
	return &UsuarioService{repository: r}
}

func (s *UsuarioService) Loggin(rs models.RequestSession) (*models.Usuario, error) {
	lenUsername := len(strings.TrimSpace(rs.Username))
	lenPassword := len(strings.TrimSpace(rs.Password))
	if lenUsername < 6 || lenPassword < 6 {
		return nil, errors.New("longitud de password o username incorrecta")
	}
	return s.repository.Loging(rs)
}
