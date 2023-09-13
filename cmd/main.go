package main

import (
	"LogApi00/database"
	"LogApi00/src/httpservice"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	app := fiber.New()
	app.Use(recover.New())
	dbCollection, err := database.Dbconnect()
	if err != nil {
		log.Fatal(err)
	}
	database.Db = dbCollection
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	httpservice.Router(app)
	appChan := make(chan error, 1)
	shutDown := make(chan os.Signal, 1)
	signal.Notify(shutDown, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		appChan <- app.Listen(":8080")
	}()
	select {
	case <-appChan:
		fmt.Printf("Listen error!\n")
	case <-shutDown:
		fmt.Printf("\nStarting Shutdown!\n")
		time.Sleep(1 * time.Second)
		fmt.Printf("Shutdown complete!\n")
	}
}
