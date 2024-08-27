package entity

import (
	"errors"
	"strconv"
	"strings"

	tele "gopkg.in/telebot.v3"
)

type User struct {
	id        int64  `db:"user_id"`
	groupSlug string `db:"group_slug"`
}

type GroupsRes struct {
	Data []string
}

type Subject struct {
	Start     string `json:"start"`
	Name      string `json:"name"`
	Teacher   string `json:"teacher"`
	Classroom string `json:"classroom"`
	Kind      Kind   `json:"kind"`
}

func (s Subject) String() string {
	var sb strings.Builder
	sb.WriteString("Начало: " + s.Start + "\n")
	sb.WriteString("Предмет: " + s.Name + "\n")
	sb.WriteString("Преподаватель: " + s.Teacher + "\n")
	return sb.String()
}

type Kind string

func (k *Kind) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), "\"")
	if str == "null" || str == "" {
		return nil
	}

	switch str {
	case "lecture":
		{
			*k = "Лекция"
			return nil
		}
	case "practice":
		{
			*k = "Практика"
			return nil
		}
	}
	return errors.New("cant unmarshal to kind type")
}

type Daily struct {
	Schedule []Subject `json:"daily_schedule"`
	Weekday  Weekday   `json:"weekday"`
}

func (d Daily) String() string {
	var sb strings.Builder
	sb.WriteString("Расписание на " + strings.ToLower(string(d.Weekday)) + "\n" + "\n")
	for _, s := range d.Schedule {
		sb.WriteString(s.String() + "\n")
	}
	return sb.String()
}

func (d *Daily) Send(b *tele.Bot, r tele.Recipient, pref *tele.SendOptions) (*tele.Message, error) {
	return b.Send(r, d.String(), pref)
}

type Weekday string

func (w *Weekday) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), "\"")
	if str == "null" || str == "" {
		return nil
	}

	switch str {

	case "monday":
		{
			*w = "Понедельник"
			return nil
		}
	case "tuesday":
		{
			*w = "Вторник"
			return nil
		}
	case "wednesday":
		{
			*w = "Среда"
			return nil
		}
	case "thursday":
		{
			*w = "Четверг"
			return nil
		}
	case "friday":
		{
			*w = "Пятница"
			return nil
		}
	case "saturday":
		{
			*w = "Суббота"
			return nil
		}
	case "sunday":
		{
			*w = "Воскресенье"
			return nil
		}
	}
	return errors.New("bad format")
}

type Weekly struct {
	Schedule []Daily `json:"weekly_schedule"`
	IsEven   bool    `json:"is_even"`
}

func (w Weekly) String() string {
	var sb strings.Builder
	sb.WriteString("Расписание на " + strconv.FormatBool(w.IsEven) + "неделю" + "\n" + "\n")
	for _, d := range w.Schedule {
		sb.WriteString(d.String())
	}
	return sb.String()
}

func (w *Weekly) Send(b *tele.Bot, r tele.Recipient, pref *tele.SendOptions) (*tele.Message, error) {
	return b.Send(r, w.String(), pref)
}
