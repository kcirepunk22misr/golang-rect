package db

import (
	"golang-react/models"
	"golang.org/x/crypto/bcrypt"
)
/*IntentoLogin*/
func IntentoLogin(email string, password string) (models.Uusario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		return usu, false
	}
	return usu, true
}

