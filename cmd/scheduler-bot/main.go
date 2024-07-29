package main

import (
	"fmt"
	"log"

	"github.com/Clankyyy/scheduler-bot/internal/pgstorage"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	storage := pgstorage.PGstorage{}
	fmt.Print(storage.GetSlug(53))
	// currenter := current.NewOSCurrenter(true)

	// token, ok := os.LookupEnv("BOT_TOKEN")
	// if !ok {
	// 	panic("Cant get bot token")
	// }

	// b := bot.NewBot(token, currenter)
	// b.Start()
}
