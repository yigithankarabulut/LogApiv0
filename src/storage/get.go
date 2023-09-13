package storage

import (
	"LogApi00/database"
	"LogApi00/src/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func GetUserByName(username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var userResult bson.M
	if err := database.Db.FindOne(ctx, bson.M{"username": username}).Decode(&userResult); err != nil {
		return err
	}
	return nil
}

func GetUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var userResult bson.M
	if err := database.Db.FindOne(ctx, bson.M{"username": user.Username, "password": user.Password}).Decode(&userResult); err != nil {
		return err
	}
	return nil
}
