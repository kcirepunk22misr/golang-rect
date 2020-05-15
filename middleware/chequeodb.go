package middleware

import (
	"net/http"
	"golang-react/db"
)
/*ChequeoDB es el middleware que me permite conocer el estado de la base de datos*/
func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if db.ChequeoConnection() == 0 {
			http.Error(writer, "Conexion perdida con la base de datos", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(writer, request)
	}
}
