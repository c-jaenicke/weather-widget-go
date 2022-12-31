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
		fmt.Fprintf(os.Stderr, "Invalid amount of arguments provided! usage: <small|full|forecast> <city name,state code,country code|.env> <path to env file>")
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
	if len(os.Args) < 4 {
		log.Fatalf("path to .env file not given")
	}

	err := godotenv.Load(os.Args[3])

	if err != nil {
		log.Fatalf("Failed to load .env file: " + err.Error())
	}

	return os.Getenv(key)
}
