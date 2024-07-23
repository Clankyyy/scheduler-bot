package main

import (
	"log"
	"os"

	"github.com/Clankyyy/scheduler-bot/internal/bot"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	token, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		panic("Cant get bot token")
	}

	b := bot.NewBot(token)
	b.Start()
}
