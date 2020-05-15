package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
/*Usuario es el modelo de usuarioa de la base de datos*/
type Uusario struct {
	ID primitive.ObjectID `bson: "_id, omitempty" json: "id"`
	Nombre string `bson: "_nombre" json: "nombre, omitempty"`
	Apellidos string `bson: "apellidos" json: "apellidos, omitempty"`
	FechaNacimiento time.Time `bson: "fechaNacimiento" json: "fechaNacimiento, omitempty"`
	Email string `bson: "email" json: "email"`
	Password string `bson: "password" json: "password, omitempty"`
	Avatar string `bson: "avatar" json: "avatar, omitempty"`
	Banner string `bson: "biografia" json: "biografia, omitempty"`
	Ubicacion string `bson: "ubicacion" json: "ubicacion, omitempty"`
	SitioWev string `bson: "sitioweb" json: "sitioweb, omitempty"`
}