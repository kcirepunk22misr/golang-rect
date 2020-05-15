package db

import (
	"context"
	"time"
	"golang-react/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertarRegistro(u models.Uusario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("golang-react")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjDB, _ := result.InsertedID.(primitive.ObjectID)
	return ObjDB.String(), true, nil
}