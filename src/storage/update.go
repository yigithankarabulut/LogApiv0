package storage

import (
	"LogApi00/database"
	"LogApi00/src/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func UpdateUser(username string, oldpwd string, newpwd string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user models.User
	user.Username = username
	user.Password = oldpwd
	if getErr := GetUser(user); getErr != nil {
		return fmt.Errorf("user with username: %s could not be found", username)
	}
	result, err := database.Db.UpdateOne(
		ctx,
		bson.M{"username": username},
		bson.D{
			{"$set", bson.D{{"password", newpwd}}},
		},
	)
	if err != nil || result.ModifiedCount == 0 {
		return fmt.Errorf("user with username: %s could not be updated", username)
	}
	return nil
}
