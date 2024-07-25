package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Clankyyy/scheduler-bot/internal/bot"
	"github.com/Clankyyy/scheduler-bot/internal/entity"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
func test() error {
	// 	data := `{"daily_schedule": [{"start": "14:00","name": "Информатика","teacher": "Федин","classroom": "416","kind": "lecture"},{"start": "15:30","name": "Русскиq","teacher": "Хз","classroom": "116","kind": "practice"}],
	//   "weekday": "thursday"}`
	some := entity.Daily{
		Schedule: []entity.Subject{
			{
				Start:     "11:30",
				Name:      "djsak",
				Teacher:   "teacher",
				Classroom: "432",
				Kind:      "ke",
			},
			{
				Start:     "14:00",
				Name:      "gagsa",
				Teacher:   "gadsg",
				Classroom: "4251",
				Kind:      "ke",
			},
		},
		Weekday: "morkov",
	}
	res, err := json.MarshalIndent(some, "", "    ")
	if err != nil {
		return err
	}
	log.Print(string(res))
	return nil
}

func main() {
	token, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		panic("Cant get bot token")
	}

	b := bot.NewBot(token)
	b.Start()
}
