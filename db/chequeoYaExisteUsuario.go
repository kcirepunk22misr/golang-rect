package db

import (
	"context"
	"time"
	"golang-react/models"
	"go.mongodb.org/mongo-driver/bson"
)
/*ChequeoYaExisteUsuario*/
func ChequeoYaExisteUsuario(email string) (models.Uusario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("golang-react")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Uusario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}