package main

import (
	"fmt"
	"log"
	"time"

	tele "gopkg.in/telebot.v4"
)

const (
	startMsg = "Привет солнышко, клацай на левый нижний угол, а там все поймешь ❤️"
	infoMsg  = `Ура солнышко! Мы вместе уже %s!
А это - %s или %s или %s вместе!
И я счаслив каждую наносекунду этого времени ❤️`
)

func main() {
	pref := tele.Settings{
		Token:     "7206554373:AAFZTIlVyTuMfo7O7aYGNZqE_LOiLiPsw5c",
		Poller:    &tele.LongPoller{Timeout: 10 * time.Second},
		ParseMode: tele.ModeHTML,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	pluralizeDays := pluralize("день", "дня", "дней")
	pluralizeHours := pluralize("час", "часа", "часов")
	pluralizeMinutes := pluralize("минута", "минуты", "минут")
	pluralizeSeconds := pluralize("секунда", "секунды", "секунд")

	zone := time.FixedZone("UTC+3", 3*60*60)
	start := time.Date(2024, time.November, 21, 9, 53, 0, 0, zone)

	b.Handle("/start", func(c tele.Context) error {
		return c.Send(startMsg)
	})

	b.Handle("/time", func(c tele.Context) error {
		now := time.Now().In(zone)
		since := now.Sub(start)
		msg := fmt.Sprintf(infoMsg,
			pluralizeDays(int(since.Hours()/24)),
			pluralizeHours(int(since.Hours())),
			pluralizeMinutes(int(since.Minutes())),
			pluralizeSeconds(int(since.Seconds())))
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
