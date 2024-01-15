package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/GeorgeHN666/werdevent-backend/app/routers"
	"github.com/GeorgeHN666/werdevent-backend/app/server"
)

func init() {

	os.Setenv("PORT", "8080")
	os.Setenv("ENV", "DEV")
	os.Setenv("VERSION", "1.0.0")
	//MAIL VARIABLES
	os.Setenv("sender-user", "contact@zkaia.com")
	os.Setenv("sender-server", "ns106.hostgator.mx")
	os.Setenv("sender-port", "465")
	os.Setenv("sender-password", "log.Fatal(1$)")

	os.Setenv("receiver-user", "")
	os.Setenv("receiver-server", "")
	os.Setenv("receiver-port", "")
	os.Setenv("receiver-password", "")

}

func main() {

	ENV := os.Getenv("ENV")

	if ENV == "DEV" {
		os.Setenv("DB", "mongodb+srv://j:rootroot@cluster0.rj0tg.mongodb.net/")
	} else if ENV == "PROD" {
		os.Setenv("DB", "PENDING")
	}

	PORT, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		fmt.Println("There was an error trying to get env variable <PORT>: ERR::: ", err.Error())
		PORT = 8080
	}

	if PORT == 0 {
		PORT = 8080
	}

	cfg := &server.Config{
		PORT:    PORT,
		Handler: routers.HandleRoutes(),
		Env:     ENV,
		Version: os.Getenv("VERSION"),
	}

	log.Fatal(server.StartServer(cfg))

}
