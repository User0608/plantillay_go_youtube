package repos

import (
	"errors"

	"github.com/user0608/api_rest/models"
	"gorm.io/gorm"
)

type UsuarioRepository struct {
	conn *gorm.DB
}

func NewUsuarioRepository(c *gorm.DB) *UsuarioRepository {
	return &UsuarioRepository{conn: c}
}

func (r *UsuarioRepository) Loging(rs models.RequestSession) (*models.Usuario, error) {
	usuario := &models.Usuario{}
	if res := r.conn.Find(usuario, "username = ? and password = ? ", rs.Username, rs.Password); res.Error != nil {
		return nil, res.Error
	} else {
		if res.RowsAffected == 0 {
			return nil, errors.New("no encotrado")
		}
	}
	return usuario, nil
}
