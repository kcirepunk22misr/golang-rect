package routers

import (
	"encoding/json"
	"net/http"
	"golang-react/db"
	"golang-react/models"
)
/*Registro es la funcion para crear la base de dato el registor de usuario*/
func Registro(writer http.ResponseWriter, r *http.Request) {

	var t models.Uusario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(writer, "Error en los datos recibidos " + err.Error(), http.StatusBadRequest)
		return
	}
	if len(t.Email) == 0 {
		http.Error(writer,"El email de usuario es requerido", http.StatusBadRequest)
		return
	}

	if len(t.Password) < 6 {
		http.Error(writer, "Debe especificar una contraseña con mas de 6 caracteres", http.StatusBadRequest)
	}

	_, encontrado, _ := db.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		 http.Error(writer, "El email ya existe", http.StatusBadRequest)
		return
	}

	_, status, err := db.InsertarRegistro(t)
	if err != nil {
		http.Error(writer, "Ocurrio un error al intentar realizar el registro de usuario" + err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(writer, "No se logro insertar el registro de usuario", http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusCreated)

}