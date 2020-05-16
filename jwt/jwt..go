package jwt

import (
	"golang-react/models"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

func GeneroJWT(t models.Uusario) (string, error)  {
	miClave := []byte("MasterDelDesarrolloGrupoDeFacebook")
	playload := jwt.MapClaims{
		"email": t.Email,
		"nombre": t.Nombre,
		"apellidos": t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia": t.Banner,
		"ubicacion": t.Ubicacion,
		"sitioweb": t.SitioWev,
		"_id": t.ID.Hex(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, playload)
	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}

