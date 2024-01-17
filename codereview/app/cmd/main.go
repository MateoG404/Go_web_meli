package main

import (
	"app/internal/application"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var (
	green = color.New(color.FgGreen).SprintFunc()
	red   = color.New(color.FgRed).SprintFunc()
)

func main() {

	// Configuration and environment variables
	// Use the environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Get the environment variables

	// - server address
	ServerAddress := os.Getenv("ADDR")

	// - loader file path

	LoaderFilePath := os.Getenv("FILEPATH")
	// app
	// - config
	cfg := &application.ConfigServerChi{
		ServerAddress:  ServerAddress,
		LoaderFilePath: LoaderFilePath,
	}

	fmt.Println(green("Running Server"))
	app := application.NewServerChi(cfg)
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(red("Error running server: "), err)
		return
	}
}
