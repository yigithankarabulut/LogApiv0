package storage

import (
	"LogApi00/database"
	"LogApi00/src/models"
	"context"
	"fmt"
	"time"
)

func CreateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := GetUserByName(user.Username); err == nil {
		return fmt.Errorf("user with username: %s already exists", user.Username)
	}
	_, err := database.Db.InsertOne(ctx, &user)
	if err != nil {
		return err
	}
	return nil
}
