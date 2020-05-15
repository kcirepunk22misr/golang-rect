package main

import (
	"log"
	"golang-react/handlers"
	"golang-react/db"
)

func main() {
	if db.ChequeoConnection()== 0 {
		log.Fatal("Sin Conexion ala db")
		return
	}
	handlers.Manejadores()
}