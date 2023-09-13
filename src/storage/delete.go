package storage

import (
	"LogApi00/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func DeleteUser(username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := GetUserByName(username); err != nil {
		return err
	}
	result, err := database.Db.DeleteOne(ctx, bson.M{"username": username})
	if err != nil || result.DeletedCount == 0 {
		return err
	}
	return nil
}
