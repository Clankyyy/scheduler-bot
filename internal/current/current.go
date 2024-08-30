package current

import (
	"log"
	"strings"
	"time"
)

type Currenter interface {
	Now() (day string, weekType string)
	NowWithOffset(days int) (string, string)
}

type OSCurrenter struct {
	shiftWeek bool
}

func NewOSCurrenter(shiftWeek bool) *OSCurrenter {
	osc := OSCurrenter{
		shiftWeek: shiftWeek,
	}

	day, kind := osc.Now()
	log.Printf("Currenter starts with day:%s, week kind: %s", day, kind)
	return &osc
}

func (osc OSCurrenter) Now() (day string, weekType string) {
	return osc.Day(0), osc.WeekKind(0)
}

func (osc OSCurrenter) NowWithOffset(days int) (string, string) {
	return osc.Day(days), osc.WeekKind(days)
}

func (osc OSCurrenter) WeekKind(offsetDays int) string {
	_, week := time.Now().AddDate(0, 0, offsetDays).ISOWeek()

	var isEven bool
	if week%2 == 0 {
		isEven = true
	} else {
		isEven = false
	}

	if osc.shiftWeek {
		isEven = !isEven
	}

	if isEven {
		return "even"
	}
	return "odd"
}

func (osc OSCurrenter) Day(offsetDays int) string {
	day := time.Now().AddDate(0, 0, offsetDays).Weekday()
	return strings.ToLower(day.String())
}
