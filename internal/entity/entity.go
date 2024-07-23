package entity

import (
	"encoding/json"
	"fmt"
)

type GroupsRes struct {
	Data []string
}

type Subject struct {
	Start     string `json:"start"`
	Name      string `json:"name"`
	Teacher   string `json:"teacher"`
	Classroom string `json:"classroom"`
	Kind      string `json:"kind"`
}

func (s *Subject) parseKind() error {
	switch s.Kind {
	case "lecture":
		s.Kind = "Лекция"
		return nil
	case "practice":
		s.Kind = "Практика"
		return nil
	}
	return fmt.Errorf("bad format")
}

type Daily struct {
	Schedule []Subject
	Weekday  string
}

func (d *Daily) UnmarshalJSON(data []byte) error {
	var daily string
	if err := json.Unmarshal(data, &daily); err != nil {
		return err
	}

	if err := d.parse(); err != nil {
		return err
	}

	return nil
}

func (d *Daily) parse() error {
	return d.parseWeekday()
}

func (d *Daily) parseWeekday() error {
	switch d.Weekday {

	case "monday":
		d.Weekday = "Понедельник"
		return nil
	case "tuesday":
		d.Weekday = "Вторник"
		return nil
	case "wednesday":
		d.Weekday = "Среда"
		return nil
	case "thursday":
		d.Weekday = "Четверг"
		return nil
	case "friday":
		d.Weekday = "Пятница"
		return nil
	case "saturday":
		d.Weekday = "Суббота"
		return nil
	case "sunday":
		d.Weekday = "Воскресенье"
		return nil
	}
	return fmt.Errorf("bad format")
}
