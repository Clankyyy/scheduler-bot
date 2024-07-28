package main

import (
	"log"
	"os"

	"github.com/Clankyyy/scheduler-bot/internal/bot"
	"github.com/Clankyyy/scheduler-bot/internal/current"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	currenter := current.NewOSCurrenter(true)
	day, kind := currenter.Now()
	log.Printf("Currenter starts with day:%s, week kind: %s", day, kind)
	token, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		panic("Cant get bot token")
	}

	b := bot.NewBot(token, currenter)
	b.Start()
}
