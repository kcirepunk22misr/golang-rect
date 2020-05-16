package handlers

import (
	"golang-react/middleware"
	"golang-react/routers"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, y pongo a escuchar el servidor*/
func Manejadores() {
	router := mux.NewRouter()
	router.HandleFunc("/registro", middleware.ChequeoDB(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middleware.ChequeoDB(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middleware.ChequeoDB(middleware.ValidoJWT(routers.VerPerfil))).Methods("GET")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}