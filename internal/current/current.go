package current

import (
	"log"
	"strings"
	"time"
)

type Currenter interface {
	Now() (string, string)
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

func (osc OSCurrenter) Now() (string, string) {
	return osc.day(), osc.weekKind()
}

func (osc OSCurrenter) weekKind() string {
	_, week := time.Now().ISOWeek()

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

func (osc OSCurrenter) day() string {
	day := time.Now().Weekday()
	return strings.ToLower(day.String())
}
