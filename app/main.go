package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/GeorgeHN666/werdevent-backend/app/router"
	"github.com/GeorgeHN666/werdevent-backend/app/server"
)

func init() {

	os.Setenv("PORT", "8000")
	os.Setenv("ENV", "DEV")
	os.Setenv("VERSION", "1.0.0")

}

func main() {

	ENV := os.Getenv("ENV")

	if ENV == "DEV" {
		os.Setenv("DB", "")
	} else if ENV == "PROD" {
		os.Setenv("DB", "")
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
		Handler: router.HandleRoutes(),
		Env:     ENV,
		Version: os.Getenv("VERSION"),
	}

	log.Fatal(server.StartServer(cfg))

}
