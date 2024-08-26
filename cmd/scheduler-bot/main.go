package main

import (
	"log"
	"os"

	"github.com/Clankyyy/scheduler-bot/internal/bot"
	"github.com/Clankyyy/scheduler-bot/internal/current"
	"github.com/Clankyyy/scheduler-bot/internal/pgstorage"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	db := pgstorage.NewPGStore()

	currenter := current.NewOSCurrenter(true)

	token, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		panic("Cant get bot token")
	}

	b := bot.NewBot(token, currenter, db)
	b.Start()
}
