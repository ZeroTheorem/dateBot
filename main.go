package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v4"
)

const (
	startMsg = "Привет солнышко, клацай на левый нижний угол, а там все поймешь ❤️"
	infoMsg  = `Ура солнышко! Мы вместе уже %s!
И я счаслив каждую наносекунду этого времени ❤️`
	infoMsg3 = `Пока что еще не прошло ни одного дня, но скоро их здесь будет бесконечное множество!
я тебя люблю ❤️`
	infoMsg4 = `Ура моя любимая жена! мы женаты 💍 уже целых %s
А впереди у нас целая жизнь и даже больше)
Бесконечность не предел ❤️
`
)

func main() {
	pref := tele.Settings{
		Token:     os.Getenv("TKN"),
		Poller:    &tele.LongPoller{Timeout: 10 * time.Second},
		ParseMode: tele.ModeHTML,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	pluralizeDays := pluralize("день", "дня", "дней")

	zone := time.FixedZone("UTC+3", 3*60*60)
	start := time.Date(2024, time.November, 21, 9, 53, 0, 0, zone)
	married := time.Date(2025, time.August, 22, 16, 30, 0, 0, zone)

	b.Handle("/start", func(c tele.Context) error {
		return c.Send(startMsg)
	})

	b.Handle("/time", func(c tele.Context) error {
		now := time.Now().In(zone)
		since := now.Sub(start)
		msg := fmt.Sprintf(infoMsg,
			pluralizeDays(int(since.Hours()/24)),
		)
		return c.Send(msg)
	})

	b.Handle("/alm", func(c tele.Context) error {
		now := time.Now().In(zone)
		since := now.Sub(married)
		switch {
		case since <= 0:
			return c.Send(infoMsg3)
		}
		msg := fmt.Sprintf(infoMsg4,
			pluralizeDays(int(since.Hours()/24)),
		)
		return c.Send(msg)
	})
	b.Start()
}

func pluralize(singular, dual, plural string) func(int) string {
	return func(n int) string {
		if n%10 == 1 && n%100 != 11 {
			return fmt.Sprintf("<b>%d</b> %s", n, singular)
		} else if (n%10 >= 2 && n%10 <= 4) && !(n%100 >= 12 && n%100 <= 14) {
			return fmt.Sprintf("<b>%d</b> %s", n, dual)
		} else {
			return fmt.Sprintf("<b>%d</b> %s", n, plural)
		}
	}
}
