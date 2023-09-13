package handlers

import (
	"fmt"
	"os"
	"time"
)

func CreateDataFile(str string, username string) error {
	newDir := fmt.Sprintf("./src/logFile/%s", username)
	if err := os.MkdirAll(newDir, 0777); err != nil {
		return err
	}
	fileName := fmt.Sprintf("./src/logFile/%s/%v.txt", username, time.Now().Unix())
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		return err
	}
	n, err := file.WriteString(str)
	if err != nil || n == 0 {
		return err
	}
	return nil
}
