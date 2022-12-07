package main

import (
	"fmt"
	"log"
	"os"
	"weather-widget-go/weather-widget-go/weather"

	"github.com/joho/godotenv"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Invalid amount of arguments provided! usage: <small|full|forecast> <city name,state code,country code> ")
		os.Exit(1)
	}

	var mode string
	var location string
	mode = args[0]
	if args[1] == ".env" {
		location = getEnv("LOCATION")
	} else {
		location = args[1]
	}

	text := weather.WeatherText(mode, location)
	fmt.Println(text)
}

func getEnv(key string) string {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("Cant load .env file")
	}

	return os.Getenv(key)
}
