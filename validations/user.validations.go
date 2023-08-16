package validations

import (
	"github.com/zeconslab/res-api-go/db"
	"github.com/zeconslab/res-api-go/models"
)

// Validar usuario existente en base de datos
func ValidarUsuario(usuario string) bool {
	var user models.User
	err := db.DB.First(&user, &usuario).Error
	if err != nil {
		return false
	}
	return true
}

// Validar email existente en base de datos
func ValidarEmail(email string) bool {
	var user models.User
	db.DB.Where("email = ?", email).First(&user)
	if user.Email == email {
		return true
	}
	return false
}
